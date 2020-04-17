// Code generated by goa v3.0.6, DO NOT EDIT.
//
// monitoring endpoints
//
// Command:
// $ goa gen github.com/anshap1719/authentication/design

package monitoring

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "monitoring" service endpoints.
type Endpoints struct {
	Status goa.Endpoint
}

// NewEndpoints wraps the methods of the "monitoring" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Status: NewStatusEndpoint(s),
	}
}

// Use applies the given middleware to all the "monitoring" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Status = m(e.Status)
}

// NewStatusEndpoint returns an endpoint function that calls the method
// "status" of service "monitoring".
func NewStatusEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, s.Status(ctx)
	}
}
