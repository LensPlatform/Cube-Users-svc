package design

import (
	. "goa.design/goa/v3/dsl"
)

// API describes the global properties of the API server.
var _ = API("Users-Microservice", server)

// Service describes a service
var _ = service
