package usersmicroserviceapi

import (
	"context"
	"fmt"
	"log"

	usersmicroservice "github.com/LensPlatform/cube_users/gen/users_microservice"
	model "github.com/LensPlatform/cube_users/pkg/models/proto"
	"goa.design/goa/v3/security"
)

// users-microservice service example implementation.
// The example methods log the requests and return zero values.
type usersMicroservicesrvc struct {
	logger *log.Logger
}

// NewUsersMicroservice returns the users-microservice service implementation.
func NewUsersMicroservice(logger *log.Logger) usersmicroservice.Service {
	return &usersMicroservicesrvc{logger}
}

// BasicAuth implements the authorization logic for service
// "users-microservice" for the "basic" security scheme.
func (s *usersMicroservicesrvc) BasicAuth(ctx context.Context, user, pass string, scheme *security.BasicScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")

	return ctx, fmt.Errorf("not implemented")
}

// Creates a valid JWT
func (s *usersMicroservicesrvc) Signin(ctx context.Context, p *usersmicroservice.SigninPayload) (res *usersmicroservice.Creds, err error) {
	res = &usersmicroservice.Creds{}
	s.logger.Print("usersMicroservice.signin")
	return
}

// CreateUser implements CreateUser.
func (s *usersMicroservicesrvc) CreateUser(ctx context.Context, p *usersmicroservice.CreateUserPayload) (res int, err error) {
	s.logger.Print("usersMicroservice.CreateUser")
	return
}

// CreateProfile implements CreateProfile.
func (s *usersMicroservicesrvc) CreateProfile(ctx context.Context, p *usersmicroservice.CreateProfilePayload) (res int, err error) {
	s.logger.Print("usersMicroservice.CreateProfile")
	return
}

// CreateUserSubscription implements CreateUserSubscription.
func (s *usersMicroservicesrvc) CreateUserSubscription(ctx context.Context, p *usersmicroservice.CreateUserSubscriptionPayload) (res int, err error) {
	s.logger.Print("usersMicroservice.CreateUserSubscription")
	return
}

// GetUser implements GetUser.
func (s *usersMicroservicesrvc) GetUser(ctx context.Context, p *usersmicroservice.GetUserPayload) (res *usersmicroservice.User, err error) {
	res = &usersmicroservice.User{}
	res.Body = model.UserORM{}
	s.logger.Print("usersMicroservice.GetUser")
	return
}
