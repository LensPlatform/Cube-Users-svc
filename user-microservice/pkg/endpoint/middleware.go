package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
)

// HiMiddleware returns an endpoint middleware
func HiMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			// Add your middleware logic here
			return next(ctx, request)
		}
	}
}
