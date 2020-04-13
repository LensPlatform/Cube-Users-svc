package design

import (
	. "goa.design/goa/v3/dsl"
)

var server = func() {
	Title("Users Microservice For The Cube Platform")
	Description("HTTP service for the users microservice")
	Docs(func() {
		Description("Service Documentation")
		URL("https://github.com/goadesign/goa/tree/master/example/security/README.md")
	})
	Server("server", func() {
		Description("server hosts the users-microservice")
		Services("users-microservice", "swagger")
		Host("development", func() {
			Description("Development Host")
			// Transport specific URLs, supported schemes are
			// http, https. grpc, grpcs with the following defaults
			// ports: 79, 443, 8080, 8443
			URI("http://localhost:7999/{version}/users-microservice")
			Variable("version", String, "API version", func() {
				Default("v1")
			})
		})

		Host("production", func() {
			Description("Productuon hosts.")
			URI("https://{version}/users-microservice")
			Variable("version", String, "API version", func() {
				Default("v1")
			})
		})
	})
}
