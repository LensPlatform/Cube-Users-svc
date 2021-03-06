// Code generated by goa v3.0.6, DO NOT EDIT.
//
// password-auth HTTP server
//
// Command:
// $ goa gen github.com/anshap1719/authentication/design

package server

import (
	"context"
	"net/http"

	passwordauth "github.com/anshap1719/authentication/controllers/gen/password_auth"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the password-auth service endpoint HTTP handlers.
type Server struct {
	Mounts              []*MountPoint
	Register            http.Handler
	Login               http.Handler
	Remove              http.Handler
	ChangePassword      http.Handler
	Reset               http.Handler
	ConfirmReset        http.Handler
	CheckEmailAvailable http.Handler
	CheckPhoneAvailable http.Handler
	CORS                http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the password-auth service endpoints.
func New(
	e *passwordauth.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Register", "POST", "/register"},
			{"Login", "POST", "/login"},
			{"Remove", "POST", "/remove-password"},
			{"ChangePassword", "POST", "/change-password"},
			{"Reset", "POST", "/reset-password"},
			{"ConfirmReset", "POST", "/finalize-reset"},
			{"CheckEmailAvailable", "POST", "/check-email-available"},
			{"CheckPhoneAvailable", "POST", "/check-phone-available"},
			{"CORS", "OPTIONS", "/register"},
			{"CORS", "OPTIONS", "/login"},
			{"CORS", "OPTIONS", "/remove-password"},
			{"CORS", "OPTIONS", "/change-password"},
			{"CORS", "OPTIONS", "/reset-password"},
			{"CORS", "OPTIONS", "/finalize-reset"},
			{"CORS", "OPTIONS", "/check-email-available"},
			{"CORS", "OPTIONS", "/check-phone-available"},
		},
		Register:            NewRegisterHandler(e.Register, mux, dec, enc, eh),
		Login:               NewLoginHandler(e.Login, mux, dec, enc, eh),
		Remove:              NewRemoveHandler(e.Remove, mux, dec, enc, eh),
		ChangePassword:      NewChangePasswordHandler(e.ChangePassword, mux, dec, enc, eh),
		Reset:               NewResetHandler(e.Reset, mux, dec, enc, eh),
		ConfirmReset:        NewConfirmResetHandler(e.ConfirmReset, mux, dec, enc, eh),
		CheckEmailAvailable: NewCheckEmailAvailableHandler(e.CheckEmailAvailable, mux, dec, enc, eh),
		CheckPhoneAvailable: NewCheckPhoneAvailableHandler(e.CheckPhoneAvailable, mux, dec, enc, eh),
		CORS:                NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "password-auth" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Register = m(s.Register)
	s.Login = m(s.Login)
	s.Remove = m(s.Remove)
	s.ChangePassword = m(s.ChangePassword)
	s.Reset = m(s.Reset)
	s.ConfirmReset = m(s.ConfirmReset)
	s.CheckEmailAvailable = m(s.CheckEmailAvailable)
	s.CheckPhoneAvailable = m(s.CheckPhoneAvailable)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the password-auth endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountRegisterHandler(mux, h.Register)
	MountLoginHandler(mux, h.Login)
	MountRemoveHandler(mux, h.Remove)
	MountChangePasswordHandler(mux, h.ChangePassword)
	MountResetHandler(mux, h.Reset)
	MountConfirmResetHandler(mux, h.ConfirmReset)
	MountCheckEmailAvailableHandler(mux, h.CheckEmailAvailable)
	MountCheckPhoneAvailableHandler(mux, h.CheckPhoneAvailable)
	MountCORSHandler(mux, h.CORS)
}

// MountRegisterHandler configures the mux to serve the "password-auth" service
// "register" endpoint.
func MountRegisterHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handlePasswordAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/register", f)
}

// NewRegisterHandler creates a HTTP handler which loads the HTTP request and
// calls the "password-auth" service "register" endpoint.
func NewRegisterHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeRegisterRequest(mux, dec)
		encodeResponse = EncodeRegisterResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "register")
		ctx = context.WithValue(ctx, goa.ServiceKey, "password-auth")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountLoginHandler configures the mux to serve the "password-auth" service
// "login" endpoint.
func MountLoginHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handlePasswordAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/login", f)
}

// NewLoginHandler creates a HTTP handler which loads the HTTP request and
// calls the "password-auth" service "login" endpoint.
func NewLoginHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeLoginRequest(mux, dec)
		encodeResponse = EncodeLoginResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "login")
		ctx = context.WithValue(ctx, goa.ServiceKey, "password-auth")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountRemoveHandler configures the mux to serve the "password-auth" service
// "remove" endpoint.
func MountRemoveHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handlePasswordAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/remove-password", f)
}

// NewRemoveHandler creates a HTTP handler which loads the HTTP request and
// calls the "password-auth" service "remove" endpoint.
func NewRemoveHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeRemoveRequest(mux, dec)
		encodeResponse = EncodeRemoveResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "remove")
		ctx = context.WithValue(ctx, goa.ServiceKey, "password-auth")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountChangePasswordHandler configures the mux to serve the "password-auth"
// service "change-password" endpoint.
func MountChangePasswordHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handlePasswordAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/change-password", f)
}

// NewChangePasswordHandler creates a HTTP handler which loads the HTTP request
// and calls the "password-auth" service "change-password" endpoint.
func NewChangePasswordHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeChangePasswordRequest(mux, dec)
		encodeResponse = EncodeChangePasswordResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "change-password")
		ctx = context.WithValue(ctx, goa.ServiceKey, "password-auth")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountResetHandler configures the mux to serve the "password-auth" service
// "reset" endpoint.
func MountResetHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handlePasswordAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/reset-password", f)
}

// NewResetHandler creates a HTTP handler which loads the HTTP request and
// calls the "password-auth" service "reset" endpoint.
func NewResetHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeResetRequest(mux, dec)
		encodeResponse = EncodeResetResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "reset")
		ctx = context.WithValue(ctx, goa.ServiceKey, "password-auth")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountConfirmResetHandler configures the mux to serve the "password-auth"
// service "confirm-reset" endpoint.
func MountConfirmResetHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handlePasswordAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/finalize-reset", f)
}

// NewConfirmResetHandler creates a HTTP handler which loads the HTTP request
// and calls the "password-auth" service "confirm-reset" endpoint.
func NewConfirmResetHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeConfirmResetRequest(mux, dec)
		encodeResponse = EncodeConfirmResetResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "confirm-reset")
		ctx = context.WithValue(ctx, goa.ServiceKey, "password-auth")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountCheckEmailAvailableHandler configures the mux to serve the
// "password-auth" service "check-email-available" endpoint.
func MountCheckEmailAvailableHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handlePasswordAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/check-email-available", f)
}

// NewCheckEmailAvailableHandler creates a HTTP handler which loads the HTTP
// request and calls the "password-auth" service "check-email-available"
// endpoint.
func NewCheckEmailAvailableHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeCheckEmailAvailableRequest(mux, dec)
		encodeResponse = EncodeCheckEmailAvailableResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "check-email-available")
		ctx = context.WithValue(ctx, goa.ServiceKey, "password-auth")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountCheckPhoneAvailableHandler configures the mux to serve the
// "password-auth" service "check-phone-available" endpoint.
func MountCheckPhoneAvailableHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handlePasswordAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/check-phone-available", f)
}

// NewCheckPhoneAvailableHandler creates a HTTP handler which loads the HTTP
// request and calls the "password-auth" service "check-phone-available"
// endpoint.
func NewCheckPhoneAvailableHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeCheckPhoneAvailableRequest(mux, dec)
		encodeResponse = EncodeCheckPhoneAvailableResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "check-phone-available")
		ctx = context.WithValue(ctx, goa.ServiceKey, "password-auth")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service password-auth.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = handlePasswordAuthOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/register", f)
	mux.Handle("OPTIONS", "/login", f)
	mux.Handle("OPTIONS", "/remove-password", f)
	mux.Handle("OPTIONS", "/change-password", f)
	mux.Handle("OPTIONS", "/reset-password", f)
	mux.Handle("OPTIONS", "/finalize-reset", f)
	mux.Handle("OPTIONS", "/check-email-available", f)
	mux.Handle("OPTIONS", "/check-phone-available", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// handlePasswordAuthOrigin applies the CORS response headers corresponding to
// the origin for the service password-auth.
func handlePasswordAuthOrigin(h http.Handler) http.Handler {
	origHndlr := h.(http.HandlerFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "*") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, X-Session")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "*")
			}
			origHndlr(w, r)
			return
		}
		origHndlr(w, r)
		return
	})
}
