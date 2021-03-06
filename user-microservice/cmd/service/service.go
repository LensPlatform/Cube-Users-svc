package service

import (
	"flag"
	"fmt"
	"net"
	http1 "net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"

	endpoint "github.com/LensPlatform/micro/user-microservice/pkg/endpoint"
	grpc "github.com/LensPlatform/micro/user-microservice/pkg/grpc"
	pb "github.com/LensPlatform/micro/user-microservice/pkg/grpc/pb"
	http "github.com/LensPlatform/micro/user-microservice/pkg/http"
	service "github.com/LensPlatform/micro/user-microservice/pkg/service"
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	lightsteptracergo "github.com/lightstep/lightstep-tracer-go"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	zipkingoopentracing "github.com/openzipkin-contrib/zipkin-go-opentracing"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	grpc1 "google.golang.org/grpc"
	appdash "sourcegraph.com/sourcegraph/appdash"
	opentracing "sourcegraph.com/sourcegraph/appdash/opentracing"
)

var tracer opentracinggo.Tracer
var logger log.Logger

var fs = flag.NewFlagSet("micro", flag.ExitOnError)
var debugAddr = fs.String("debug.addr", ":8086", "Debug and metrics listen address")
var httpAddr = fs.String("http-addr", ":8087", "HTTP listen address")
var grpcAddr = fs.String("grpc-addr", ":8088", "gRPC listen address")
var thriftAddr = fs.String("thrift-addr", ":8089", "Thrift listen address")
var thriftProtocol = fs.String("thrift-protocol", "binary", "binary, compact, json, simplejson")
var thriftBuffer = fs.Int("thrift-buffer", 0, "0 for unbuffered")
var thriftFramed = fs.Bool("thrift-framed", false, "true to enable framing")
var zipkinURL = fs.String("zipkin-url", "", "Enable Zipkin tracing via a collector URL e.g. http://localhost:9411/api/v1/spans")
var lightstepToken = fs.String("lightstep-token", "", "Enable LightStep tracing via a LightStep access token")
var appdashAddr = fs.String("appdash-addr", "", "Enable Appdash tracing via an Appdash server host:port")
var connType = fs.String("db-conn-type", "postgresql://", "database connection type")
var connAddress = fs.String("db-conn-address", "doadmin:oqshd3sto72yyhgq@test-do-user-6612421-0.a.db.ondigitalocean.com:25060/", "database connection address")
var connName = fs.String("db-conn-name", "defaultdb", "Database connection name")
var connSettings = fs.String("db-conn-settings", "?sslmode=require", "database connection settings")

func Run() {
	fs.Parse(os.Args[1:])

	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	connectionString := *connType + *connAddress + *connName + *connSettings

	if *zipkinURL != "" {
		/*
			logger.Log("tracer", "Zipkin", "URL", *zipkinURL)
			collector, err := zipkingoopentracing.NewHTTPCollector(*zipkinURL)
			if err != nil {
				logger.Log("err", err)
				os.Exit(1)
			}
			defer collector.Close()
			recorder := zipkingoopentracing.NewRecorder(collector, false, "localhost:80", "micro")
			tracer, err = zipkingoopentracing.NewTracer(recorder)
			if err != nil {
				logger.Log("err", err)
				os.Exit(1)
			}

			// set up a span reporter
			reporter := zipkinhttp.NewReporter("http://zipkinhost:9411/api/v2/spans")
			defer reporter.Close()

			// create our local service endpoint
			endpoint, err := zipkin.NewEndpoint("myService", "myservice.mydomain.com:80")
			if err != nil {
				log.Fatalf("unable to create local endpoint: %+v\n", err)
			}
		*/
		// set up a span reporter
		reporter := zipkinhttp.NewReporter("http://zipkinhost:9411/api/v2/spans")
		defer reporter.Close()

		// create our local service endpoint
		endpoint, err := zipkin.NewEndpoint("Cube_Users-Microservice", "localhost:80")
		if err != nil {
			logger.Log("unable to create tracer: %+v\n", err)
		}
		// initialize our tracer
		nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
		if err != nil {
			logger.Log("unable to create tracer: %+v\n", err)
		}

		// use zipkin-go-opentracing to wrap our tracer
		tracer := zipkingoopentracing.Wrap(nativeTracer)

		// optionally set as Global OpenTracing tracer instance
		opentracinggo.SetGlobalTracer(tracer)
	} else if *lightstepToken != "" {
		logger.Log("tracer", "LightStep")
		tracer = lightsteptracergo.NewTracer(lightsteptracergo.Options{AccessToken: *lightstepToken})
		defer lightsteptracergo.FlushLightStepTracer(tracer)
	} else if *appdashAddr != "" {
		logger.Log("tracer", "Appdash", "addr", *appdashAddr)
		collector := appdash.NewRemoteCollector(*appdashAddr)
		tracer = opentracing.NewTracer(collector)
		defer collector.Close()
	} else {
		logger.Log("tracer", "none")
		tracer = opentracinggo.GlobalTracer()
	}

	svc := service.New(getServiceMiddleware(logger), connectionString, logger)
	eps := endpoint.New(svc, getEndpointMiddleware(logger))
	g := createService(eps)
	initMetricsEndpoint(g)
	initCancelInterrupt(g)
	logger.Log("exit", g.Run())

}
func initHttpHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultHttpOptions(logger, tracer)

	httpHandler := http.NewHTTPHandler(endpoints, options)
	httpListener, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		logger.Log("transport", "HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		return http1.Serve(httpListener, httpHandler)
	}, func(error) {
		httpListener.Close()
	})

}
func getServiceMiddleware(logger log.Logger) (mw []service.Middleware) {
	mw = []service.Middleware{}

	return
}
func getEndpointMiddleware(logger log.Logger) (mw map[string][]endpoint1.Middleware) {
	mw = map[string][]endpoint1.Middleware{}

	return
}
func initMetricsEndpoint(g *group.Group) {
	http1.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	debugListener, err := net.Listen("tcp", *debugAddr)
	if err != nil {
		logger.Log("transport", "debug/HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "debug/HTTP", "addr", *debugAddr)
		return http1.Serve(debugListener, http1.DefaultServeMux)
	}, func(error) {
		debugListener.Close()
	})
}
func initCancelInterrupt(g *group.Group) {
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
}

func initGRPCHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultGRPCOptions(logger, tracer)

	grpcServer := grpc.NewGRPCServer(endpoints, options)
	grpcListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "gRPC", "addr", *grpcAddr)
		baseServer := grpc1.NewServer()
		pb.RegisterMicroServer(baseServer, grpcServer)
		return baseServer.Serve(grpcListener)
	}, func(error) {
		grpcListener.Close()
	})

}
