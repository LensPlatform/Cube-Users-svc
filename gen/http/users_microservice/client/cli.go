// Code generated by goa v3.1.1, DO NOT EDIT.
//
// users-microservice HTTP client CLI support package
//
// Command:
// $ goa gen github.com/LensPlatform/cube_users/design

package client

import (
	"encoding/json"
	"fmt"

	usersmicroservice "github.com/LensPlatform/cube_users/gen/users_microservice"
)

// BuildSigninPayload builds the payload for the users-microservice signin
// endpoint from CLI flags.
func BuildSigninPayload(usersMicroserviceSigninBody string, usersMicroserviceSigninUsername string, usersMicroserviceSigninPassword string) (*usersmicroservice.SigninPayload, error) {
	var err error
	var body SigninRequestBody
	{
		err = json.Unmarshal([]byte(usersMicroserviceSigninBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"email\": \"userame@gmail.com\"\n   }'")
		}
	}
	var username string
	{
		username = usersMicroserviceSigninUsername
	}
	var password string
	{
		password = usersMicroserviceSigninPassword
	}
	v := &usersmicroservice.SigninPayload{
		Email: body.Email,
	}
	v.Username = username
	v.Password = password

	return v, nil
}

// BuildCreateUserPayload builds the payload for the users-microservice
// CreateUser endpoint from CLI flags.
func BuildCreateUserPayload(usersMicroserviceCreateUserBody string) (*usersmicroservice.CreateUserPayload, error) {
	var err error
	var body CreateUserRequestBody
	{
		err = json.Unmarshal([]byte(usersMicroserviceCreateUserBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"user\": \"Reprehenderit illo.\"\n   }'")
		}
	}
	v := &usersmicroservice.CreateUserPayload{
		User: body.User,
	}

	return v, nil
}

// BuildCreateProfilePayload builds the payload for the users-microservice
// CreateProfile endpoint from CLI flags.
func BuildCreateProfilePayload(usersMicroserviceCreateProfileBody string) (*usersmicroservice.CreateProfilePayload, error) {
	var err error
	var body CreateProfileRequestBody
	{
		err = json.Unmarshal([]byte(usersMicroserviceCreateProfileBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"profile\": \"Ab magni consequatur tenetur cupiditate.\",\n      \"user_id\": \"Voluptatem enim neque.\"\n   }'")
		}
	}
	v := &usersmicroservice.CreateProfilePayload{
		Profile: body.Profile,
		UserID:  body.UserID,
	}

	return v, nil
}

// BuildCreateUserSubscriptionPayload builds the payload for the
// users-microservice CreateUserSubscription endpoint from CLI flags.
func BuildCreateUserSubscriptionPayload(usersMicroserviceCreateUserSubscriptionBody string) (*usersmicroservice.CreateUserSubscriptionPayload, error) {
	var err error
	var body CreateUserSubscriptionRequestBody
	{
		err = json.Unmarshal([]byte(usersMicroserviceCreateUserSubscriptionBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"subscription\": \"Quae est.\",\n      \"user_id\": \"Deserunt error ipsum quas eum rerum repudiandae.\"\n   }'")
		}
	}
	v := &usersmicroservice.CreateUserSubscriptionPayload{
		Subscription: body.Subscription,
		UserID:       body.UserID,
	}

	return v, nil
}

// BuildGetUserPayload builds the payload for the users-microservice GetUser
// endpoint from CLI flags.
func BuildGetUserPayload(usersMicroserviceGetUserUserID string) (*usersmicroservice.GetUserPayload, error) {
	var userID string
	{
		userID = usersMicroserviceGetUserUserID
	}
	v := &usersmicroservice.GetUserPayload{}
	v.UserID = userID

	return v, nil
}
