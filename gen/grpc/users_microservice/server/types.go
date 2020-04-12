// Code generated by goa v3.1.1, DO NOT EDIT.
//
// users-microservice gRPC server types
//
// Command:
// $ goa gen github.com/LensPlatform/cube_users/design

package server

import (
	users_microservicepb "github.com/LensPlatform/cube_users/gen/grpc/users_microservice/pb"
	usersmicroservice "github.com/LensPlatform/cube_users/gen/users_microservice"
)

// NewSigninPayload builds the payload of the "signin" endpoint of the
// "users-microservice" service from the gRPC request type.
func NewSigninPayload(message *users_microservicepb.SigninRequest, username string, password string) *usersmicroservice.SigninPayload {
	v := &usersmicroservice.SigninPayload{}
	if message.Email != "" {
		v.Email = &message.Email
	}
	v.Username = username
	v.Password = password
	return v
}

// NewSigninResponse builds the gRPC response type from the result of the
// "signin" endpoint of the "users-microservice" service.
func NewSigninResponse(result *usersmicroservice.Creds) *users_microservicepb.SigninResponse {
	message := &users_microservicepb.SigninResponse{
		Jwt:        result.JWT,
		ApiKey:     result.APIKey,
		OauthToken: result.OauthToken,
	}
	return message
}

// NewCreateUserPayload builds the payload of the "CreateUser" endpoint of the
// "users-microservice" service from the gRPC request type.
func NewCreateUserPayload(message *users_microservicepb.CreateUserRequest) *usersmicroservice.CreateUserPayload {
	v := &usersmicroservice.CreateUserPayload{
		User: message.User,
	}
	return v
}

// NewCreateUserResponse builds the gRPC response type from the result of the
// "CreateUser" endpoint of the "users-microservice" service.
func NewCreateUserResponse(result string) *users_microservicepb.CreateUserResponse {
	message := &users_microservicepb.CreateUserResponse{}
	message.Field = result
	return message
}

// NewCreateProfilePayload builds the payload of the "CreateProfile" endpoint
// of the "users-microservice" service from the gRPC request type.
func NewCreateProfilePayload(message *users_microservicepb.CreateProfileRequest) *usersmicroservice.CreateProfilePayload {
	v := &usersmicroservice.CreateProfilePayload{
		Profile: message.Profile,
		UserID:  message.UserId,
	}
	return v
}

// NewCreateProfileResponse builds the gRPC response type from the result of
// the "CreateProfile" endpoint of the "users-microservice" service.
func NewCreateProfileResponse(result string) *users_microservicepb.CreateProfileResponse {
	message := &users_microservicepb.CreateProfileResponse{}
	message.Field = result
	return message
}

// NewCreateUserSubscriptionPayload builds the payload of the
// "CreateUserSubscription" endpoint of the "users-microservice" service from
// the gRPC request type.
func NewCreateUserSubscriptionPayload(message *users_microservicepb.CreateUserSubscriptionRequest) *usersmicroservice.CreateUserSubscriptionPayload {
	v := &usersmicroservice.CreateUserSubscriptionPayload{
		Subscription: message.Subscription,
		UserID:       message.UserId,
	}
	return v
}

// NewCreateUserSubscriptionResponse builds the gRPC response type from the
// result of the "CreateUserSubscription" endpoint of the "users-microservice"
// service.
func NewCreateUserSubscriptionResponse(result string) *users_microservicepb.CreateUserSubscriptionResponse {
	message := &users_microservicepb.CreateUserSubscriptionResponse{}
	message.Field = result
	return message
}

// NewGetUserPayload builds the payload of the "GetUser" endpoint of the
// "users-microservice" service from the gRPC request type.
func NewGetUserPayload(message *users_microservicepb.GetUserRequest) *usersmicroservice.GetUserPayload {
	v := &usersmicroservice.GetUserPayload{
		UserID: message.UserId,
	}
	return v
}

// NewGetUserResponse builds the gRPC response type from the result of the
// "GetUser" endpoint of the "users-microservice" service.
func NewGetUserResponse(result string) *users_microservicepb.GetUserResponse {
	message := &users_microservicepb.GetUserResponse{}
	message.Field = result
	return message
}
