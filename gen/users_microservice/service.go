// Code generated by goa v3.1.1, DO NOT EDIT.
//
// users-microservice service
//
// Command:
// $ goa gen github.com/LensPlatform/cube_users/design

package usersmicroservice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// The users microservice exposes endpoStrings useful in accessing various
// schema types
type Service interface {
	// Creates a valid JWT
	Signin(context.Context, *SigninPayload) (res *Creds, err error)
	// CreateUser implements CreateUser.
	CreateUser(context.Context, *CreateUserPayload) (res int, err error)
	// CreateProfile implements CreateProfile.
	CreateProfile(context.Context, *CreateProfilePayload) (res int, err error)
	// CreateUserSubscription implements CreateUserSubscription.
	CreateUserSubscription(context.Context, *CreateUserSubscriptionPayload) (res int, err error)
	// GetUser implements GetUser.
	GetUser(context.Context, *GetUserPayload) (res *User, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// BasicAuth implements the authorization logic for the Basic security scheme.
	BasicAuth(ctx context.Context, user, pass string, schema *security.BasicScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "users-microservice"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"signin", "CreateUser", "CreateProfile", "CreateUserSubscription", "GetUser"}

// Credentials used to authenticate to retrieve JWT token
type SigninPayload struct {
	// Username used to perform signin
	Username string
	// Password used to perform signin
	Password string
	// Email used to perform sigin
	Email *string
}

// Creds is the result type of the users-microservice service signin method.
type Creds struct {
	// JWT token
	JWT string
	// API Key
	APIKey string
	// OAuth2 token
	OauthToken string
}

// CreateUserPayload is the payload type of the users-microservice service
// CreateUser method.
type CreateUserPayload struct {
	// User to be created
	User *User
}

// CreateProfilePayload is the payload type of the users-microservice service
// CreateProfile method.
type CreateProfilePayload struct {
	// Profile
	Profile *Profile
	// user id token which the profile is tied to
	UserID int64
}

// CreateUserSubscriptionPayload is the payload type of the users-microservice
// service CreateUserSubscription method.
type CreateUserSubscriptionPayload struct {
	// User Subscription
	Subscription *Subscription
	// user id to which the subscription is to be created for
	UserID int64
}

// GetUserPayload is the payload type of the users-microservice service GetUser
// method.
type GetUserPayload struct {
	// User id
	UserID int64
}

// User is the result type of the users-microservice service GetUser method.
type User struct {
	Body interface{}
}

type Profile struct {
	Body interface{}
}

type Subscription struct {
	Body interface{}
}

// Credentials are invalid
type Unauthorized string

// Error returns an error description.
func (e Unauthorized) Error() string {
	return "Credentials are invalid"
}

// ErrorName returns "unauthorized".
func (e Unauthorized) ErrorName() string {
	return "unauthorized"
}

// MakeTimeout builds a goa.ServiceError from an error.
func MakeTimeout(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:      "timeout",
		ID:        goa.NewErrorID(),
		Message:   err.Error(),
		Temporary: true,
		Timeout:   true,
	}
}
