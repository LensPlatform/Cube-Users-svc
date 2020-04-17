package http

import (
	"context"
	"encoding/json"
	"errors"
	endpoint "github.com/LensPlatform/micro/user-microservice/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	"net/http"
)

// makeCreateUserHandler creates the handler logic
func makeCreateUserHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-user", http1.NewServer(endpoints.CreateUserEndpoint, decodeCreateUserRequest, encodeCreateUserResponse, options...))
}

// decodeCreateUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateUserResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateProfileHandler creates the handler logic
func makeCreateProfileHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-profile", http1.NewServer(endpoints.CreateProfileEndpoint, decodeCreateProfileRequest, encodeCreateProfileResponse, options...))
}

// decodeCreateProfileRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateProfileRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateProfileResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateProfileResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateSubscriptionHandler creates the handler logic
func makeCreateSubscriptionHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-subscription", http1.NewServer(endpoints.CreateSubscriptionEndpoint, decodeCreateSubscriptionRequest, encodeCreateSubscriptionResponse, options...))
}

// decodeCreateSubscriptionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateSubscriptionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateSubscriptionRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateSubscriptionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateSubscriptionsHandler creates the handler logic
func makeCreateSubscriptionsHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-subscriptions", http1.NewServer(endpoints.CreateSubscriptionsEndpoint, decodeCreateSubscriptionsRequest, encodeCreateSubscriptionsResponse, options...))
}

// decodeCreateSubscriptionsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateSubscriptionsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateSubscriptionsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateSubscriptionsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateSubscriptionsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateUserHandler creates the handler logic
func makeUpdateUserHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-user", http1.NewServer(endpoints.UpdateUserEndpoint, decodeUpdateUserRequest, encodeUpdateUserResponse, options...))
}

// decodeUpdateUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateUserResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateUserProfileHandler creates the handler logic
func makeUpdateUserProfileHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-user-profile", http1.NewServer(endpoints.UpdateUserProfileEndpoint, decodeUpdateUserProfileRequest, encodeUpdateUserProfileResponse, options...))
}

// decodeUpdateUserProfileRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateUserProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateUserProfileRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateUserProfileResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateUserProfileResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateUserSubscriptionHandler creates the handler logic
func makeUpdateUserSubscriptionHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-user-subscription", http1.NewServer(endpoints.UpdateUserSubscriptionEndpoint, decodeUpdateUserSubscriptionRequest, encodeUpdateUserSubscriptionResponse, options...))
}

// decodeUpdateUserSubscriptionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateUserSubscriptionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateUserSubscriptionRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateUserSubscriptionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateUserSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteUserHandler creates the handler logic
func makeDeleteUserHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-user", http1.NewServer(endpoints.DeleteUserEndpoint, decodeDeleteUserRequest, encodeDeleteUserResponse, options...))
}

// decodeDeleteUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteUserResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteUserProfileHandler creates the handler logic
func makeDeleteUserProfileHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-user-profile", http1.NewServer(endpoints.DeleteUserProfileEndpoint, decodeDeleteUserProfileRequest, encodeDeleteUserProfileResponse, options...))
}

// decodeDeleteUserProfileRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteUserProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteUserProfileRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteUserProfileResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteUserProfileResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteSubscriptionHandler creates the handler logic
func makeDeleteSubscriptionHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-subscription", http1.NewServer(endpoints.DeleteSubscriptionEndpoint, decodeDeleteSubscriptionRequest, encodeDeleteSubscriptionResponse, options...))
}

// decodeDeleteSubscriptionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteSubscriptionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteSubscriptionRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteSubscriptionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetUserHandler creates the handler logic
func makeGetUserHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-user", http1.NewServer(endpoints.GetUserEndpoint, decodeGetUserRequest, encodeGetUserResponse, options...))
}

// decodeGetUserRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetUserResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetUsersHandler creates the handler logic
func makeGetUsersHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-users", http1.NewServer(endpoints.GetUsersEndpoint, decodeGetUsersRequest, encodeGetUsersResponse, options...))
}

// decodeGetUsersRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetUsersRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetUsersResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetUsersResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetUsersByAccountTypeHandler creates the handler logic
func makeGetUsersByAccountTypeHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-users-by-account-type", http1.NewServer(endpoints.GetUsersByAccountTypeEndpoint, decodeGetUsersByAccountTypeRequest, encodeGetUsersByAccountTypeResponse, options...))
}

// decodeGetUsersByAccountTypeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetUsersByAccountTypeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetUsersByAccountTypeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetUsersByAccountTypeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetUsersByAccountTypeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetUsersByIntentHandler creates the handler logic
func makeGetUsersByIntentHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-users-by-intent", http1.NewServer(endpoints.GetUsersByIntentEndpoint, decodeGetUsersByIntentRequest, encodeGetUsersByIntentResponse, options...))
}

// decodeGetUsersByIntentRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetUsersByIntentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetUsersByIntentRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetUsersByIntentResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetUsersByIntentResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetUserProfileHandler creates the handler logic
func makeGetUserProfileHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-user-profile", http1.NewServer(endpoints.GetUserProfileEndpoint, decodeGetUserProfileRequest, encodeGetUserProfileResponse, options...))
}

// decodeGetUserProfileRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetUserProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetUserProfileRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetUserProfileResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetUserProfileResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetUserProfilesHandler creates the handler logic
func makeGetUserProfilesHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-user-profiles", http1.NewServer(endpoints.GetUserProfilesEndpoint, decodeGetUserProfilesRequest, encodeGetUserProfilesResponse, options...))
}

// decodeGetUserProfilesRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetUserProfilesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetUserProfilesRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetUserProfilesResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetUserProfilesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetUserProfilesByTypeHandler creates the handler logic
func makeGetUserProfilesByTypeHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-user-profiles-by-type", http1.NewServer(endpoints.GetUserProfilesByTypeEndpoint, decodeGetUserProfilesByTypeRequest, encodeGetUserProfilesByTypeResponse, options...))
}

// decodeGetUserProfilesByTypeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetUserProfilesByTypeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetUserProfilesByTypeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetUserProfilesByTypeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetUserProfilesByTypeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetUserProfilesByNationalityHandler creates the handler logic
func makeGetUserProfilesByNationalityHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-user-profiles-by-nationality", http1.NewServer(endpoints.GetUserProfilesByNationalityEndpoint, decodeGetUserProfilesByNationalityRequest, encodeGetUserProfilesByNationalityResponse, options...))
}

// decodeGetUserProfilesByNationalityRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetUserProfilesByNationalityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetUserProfilesByNationalityRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetUserProfilesByNationalityResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetUserProfilesByNationalityResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetUserSubscriptionsHandler creates the handler logic
func makeGetUserSubscriptionsHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-user-subscriptions", http1.NewServer(endpoints.GetUserSubscriptionsEndpoint, decodeGetUserSubscriptionsRequest, encodeGetUserSubscriptionsResponse, options...))
}

// decodeGetUserSubscriptionsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetUserSubscriptionsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetUserSubscriptionsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetUserSubscriptionsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetUserSubscriptionsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateTeamHandler creates the handler logic
func makeCreateTeamHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-team", http1.NewServer(endpoints.CreateTeamEndpoint, decodeCreateTeamRequest, encodeCreateTeamResponse, options...))
}

// decodeCreateTeamRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateTeamRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateTeamRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateTeamResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateTeamResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreatTeamProfileHandler creates the handler logic
func makeCreatTeamProfileHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/creat-team-profile", http1.NewServer(endpoints.CreatTeamProfileEndpoint, decodeCreatTeamProfileRequest, encodeCreatTeamProfileResponse, options...))
}

// decodeCreatTeamProfileRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreatTeamProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreatTeamProfileRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreatTeamProfileResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreatTeamProfileResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateTeamSubscriptionHandler creates the handler logic
func makeCreateTeamSubscriptionHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-team-subscription", http1.NewServer(endpoints.CreateTeamSubscriptionEndpoint, decodeCreateTeamSubscriptionRequest, encodeCreateTeamSubscriptionResponse, options...))
}

// decodeCreateTeamSubscriptionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateTeamSubscriptionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateTeamSubscriptionRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateTeamSubscriptionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateTeamSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAddMemberToTeamHandler creates the handler logic
func makeAddMemberToTeamHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/add-member-to-team", http1.NewServer(endpoints.AddMemberToTeamEndpoint, decodeAddMemberToTeamRequest, encodeAddMemberToTeamResponse, options...))
}

// decodeAddMemberToTeamRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddMemberToTeamRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.AddMemberToTeamRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeAddMemberToTeamResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddMemberToTeamResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAddAdvisorToTeamHandler creates the handler logic
func makeAddAdvisorToTeamHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/add-advisor-to-team", http1.NewServer(endpoints.AddAdvisorToTeamEndpoint, decodeAddAdvisorToTeamRequest, encodeAddAdvisorToTeamResponse, options...))
}

// decodeAddAdvisorToTeamRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddAdvisorToTeamRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.AddAdvisorToTeamRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeAddAdvisorToTeamResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddAdvisorToTeamResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateTeamHandler creates the handler logic
func makeUpdateTeamHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-team", http1.NewServer(endpoints.UpdateTeamEndpoint, decodeUpdateTeamRequest, encodeUpdateTeamResponse, options...))
}

// decodeUpdateTeamRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateTeamRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateTeamRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateTeamResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateTeamResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateTeamAdminHandler creates the handler logic
func makeUpdateTeamAdminHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-team-admin", http1.NewServer(endpoints.UpdateTeamAdminEndpoint, decodeUpdateTeamAdminRequest, encodeUpdateTeamAdminResponse, options...))
}

// decodeUpdateTeamAdminRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateTeamAdminRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateTeamAdminRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateTeamAdminResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateTeamAdminResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateTeamMemberHandler creates the handler logic
func makeUpdateTeamMemberHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-team-member", http1.NewServer(endpoints.UpdateTeamMemberEndpoint, decodeUpdateTeamMemberRequest, encodeUpdateTeamMemberResponse, options...))
}

// decodeUpdateTeamMemberRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateTeamMemberRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateTeamMemberRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateTeamMemberResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateTeamMemberResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateTeamAdvisorHandler creates the handler logic
func makeUpdateTeamAdvisorHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-team-advisor", http1.NewServer(endpoints.UpdateTeamAdvisorEndpoint, decodeUpdateTeamAdvisorRequest, encodeUpdateTeamAdvisorResponse, options...))
}

// decodeUpdateTeamAdvisorRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateTeamAdvisorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateTeamAdvisorRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateTeamAdvisorResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateTeamAdvisorResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateTeamProfileHandler creates the handler logic
func makeUpdateTeamProfileHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-team-profile", http1.NewServer(endpoints.UpdateTeamProfileEndpoint, decodeUpdateTeamProfileRequest, encodeUpdateTeamProfileResponse, options...))
}

// decodeUpdateTeamProfileRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateTeamProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateTeamProfileRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateTeamProfileResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateTeamProfileResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateTeamSubscriptionHandler creates the handler logic
func makeUpdateTeamSubscriptionHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-team-subscription", http1.NewServer(endpoints.UpdateTeamSubscriptionEndpoint, decodeUpdateTeamSubscriptionRequest, encodeUpdateTeamSubscriptionResponse, options...))
}

// decodeUpdateTeamSubscriptionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateTeamSubscriptionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateTeamSubscriptionRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateTeamSubscriptionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateTeamSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteTeamHandler creates the handler logic
func makeDeleteTeamHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-team", http1.NewServer(endpoints.DeleteTeamEndpoint, decodeDeleteTeamRequest, encodeDeleteTeamResponse, options...))
}

// decodeDeleteTeamRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteTeamRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteTeamRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteTeamResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteTeamResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteTeamProfileHandler creates the handler logic
func makeDeleteTeamProfileHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-team-profile", http1.NewServer(endpoints.DeleteTeamProfileEndpoint, decodeDeleteTeamProfileRequest, encodeDeleteTeamProfileResponse, options...))
}

// decodeDeleteTeamProfileRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteTeamProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteTeamProfileRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteTeamProfileResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteTeamProfileResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteTeamMemberHandler creates the handler logic
func makeDeleteTeamMemberHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-team-member", http1.NewServer(endpoints.DeleteTeamMemberEndpoint, decodeDeleteTeamMemberRequest, encodeDeleteTeamMemberResponse, options...))
}

// decodeDeleteTeamMemberRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteTeamMemberRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteTeamMemberRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteTeamMemberResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteTeamMemberResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteTeamAdvisorHandler creates the handler logic
func makeDeleteTeamAdvisorHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-team-advisor", http1.NewServer(endpoints.DeleteTeamAdvisorEndpoint, decodeDeleteTeamAdvisorRequest, encodeDeleteTeamAdvisorResponse, options...))
}

// decodeDeleteTeamAdvisorRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteTeamAdvisorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteTeamAdvisorRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteTeamAdvisorResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteTeamAdvisorResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteTeamAdminHandler creates the handler logic
func makeDeleteTeamAdminHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-team-admin", http1.NewServer(endpoints.DeleteTeamAdminEndpoint, decodeDeleteTeamAdminRequest, encodeDeleteTeamAdminResponse, options...))
}

// decodeDeleteTeamAdminRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteTeamAdminRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteTeamAdminRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteTeamAdminResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteTeamAdminResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteTeamSubscriptionHandler creates the handler logic
func makeDeleteTeamSubscriptionHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-team-subscription", http1.NewServer(endpoints.DeleteTeamSubscriptionEndpoint, decodeDeleteTeamSubscriptionRequest, encodeDeleteTeamSubscriptionResponse, options...))
}

// decodeDeleteTeamSubscriptionRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteTeamSubscriptionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteTeamSubscriptionRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteTeamSubscriptionResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteTeamSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetTeamHandler creates the handler logic
func makeGetTeamHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-team", http1.NewServer(endpoints.GetTeamEndpoint, decodeGetTeamRequest, encodeGetTeamResponse, options...))
}

// decodeGetTeamRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetTeamRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetTeamRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetTeamResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetTeamResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetTeamProfileHandler creates the handler logic
func makeGetTeamProfileHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-team-profile", http1.NewServer(endpoints.GetTeamProfileEndpoint, decodeGetTeamProfileRequest, encodeGetTeamProfileResponse, options...))
}

// decodeGetTeamProfileRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetTeamProfileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetTeamProfileRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetTeamProfileResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetTeamProfileResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetTeamSubscriptionsHandler creates the handler logic
func makeGetTeamSubscriptionsHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-team-subscriptions", http1.NewServer(endpoints.GetTeamSubscriptionsEndpoint, decodeGetTeamSubscriptionsRequest, encodeGetTeamSubscriptionsResponse, options...))
}

// decodeGetTeamSubscriptionsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetTeamSubscriptionsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetTeamSubscriptionsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetTeamSubscriptionsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetTeamSubscriptionsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetTeamMembersHandler creates the handler logic
func makeGetTeamMembersHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-team-members", http1.NewServer(endpoints.GetTeamMembersEndpoint, decodeGetTeamMembersRequest, encodeGetTeamMembersResponse, options...))
}

// decodeGetTeamMembersRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetTeamMembersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetTeamMembersRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetTeamMembersResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetTeamMembersResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetTeamAdvisorsHandler creates the handler logic
func makeGetTeamAdvisorsHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-team-advisors", http1.NewServer(endpoints.GetTeamAdvisorsEndpoint, decodeGetTeamAdvisorsRequest, encodeGetTeamAdvisorsResponse, options...))
}

// decodeGetTeamAdvisorsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetTeamAdvisorsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetTeamAdvisorsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetTeamAdvisorsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetTeamAdvisorsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetTeamAdminHandler creates the handler logic
func makeGetTeamAdminHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-team-admin", http1.NewServer(endpoints.GetTeamAdminEndpoint, decodeGetTeamAdminRequest, encodeGetTeamAdminResponse, options...))
}

// decodeGetTeamAdminRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetTeamAdminRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetTeamAdminRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetTeamAdminResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetTeamAdminResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetTeamsHandler creates the handler logic
func makeGetTeamsHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-teams", http1.NewServer(endpoints.GetTeamsEndpoint, decodeGetTeamsRequest, encodeGetTeamsResponse, options...))
}

// decodeGetTeamsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetTeamsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetTeamsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetTeamsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetTeamsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetTeamsByTypeHandler creates the handler logic
func makeGetTeamsByTypeHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-teams-by-type", http1.NewServer(endpoints.GetTeamsByTypeEndpoint, decodeGetTeamsByTypeRequest, encodeGetTeamsByTypeResponse, options...))
}

// decodeGetTeamsByTypeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetTeamsByTypeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetTeamsByTypeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetTeamsByTypeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetTeamsByTypeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetTeamsByIndustryHandler creates the handler logic
func makeGetTeamsByIndustryHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-teams-by-industry", http1.NewServer(endpoints.GetTeamsByIndustryEndpoint, decodeGetTeamsByIndustryRequest, encodeGetTeamsByIndustryResponse, options...))
}

// decodeGetTeamsByIndustryRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetTeamsByIndustryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetTeamsByIndustryRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetTeamsByIndustryResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetTeamsByIndustryResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetTeamsByNumberOfEmployeesHandler creates the handler logic
func makeGetTeamsByNumberOfEmployeesHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-teams-by-number-of-employees", http1.NewServer(endpoints.GetTeamsByNumberOfEmployeesEndpoint, decodeGetTeamsByNumberOfEmployeesRequest, encodeGetTeamsByNumberOfEmployeesResponse, options...))
}

// decodeGetTeamsByNumberOfEmployeesRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetTeamsByNumberOfEmployeesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetTeamsByNumberOfEmployeesRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetTeamsByNumberOfEmployeesResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetTeamsByNumberOfEmployeesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetTeamProfilesHandler creates the handler logic
func makeGetTeamProfilesHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-team-profiles", http1.NewServer(endpoints.GetTeamProfilesEndpoint, decodeGetTeamProfilesRequest, encodeGetTeamProfilesResponse, options...))
}

// decodeGetTeamProfilesRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetTeamProfilesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetTeamProfilesRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetTeamProfilesResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetTeamProfilesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetTeamsByTagsHandler creates the handler logic
func makeGetTeamsByTagsHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-teams-by-tags", http1.NewServer(endpoints.GetTeamsByTagsEndpoint, decodeGetTeamsByTagsRequest, encodeGetTeamsByTagsResponse, options...))
}

// decodeGetTeamsByTagsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetTeamsByTagsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetTeamsByTagsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetTeamsByTagsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetTeamsByTagsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateGroupHandler creates the handler logic
func makeCreateGroupHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-group", http1.NewServer(endpoints.CreateGroupEndpoint, decodeCreateGroupRequest, encodeCreateGroupResponse, options...))
}

// decodeCreateGroupRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateGroupRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateGroupResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateGroupResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAddGroupMemberHandler creates the handler logic
func makeAddGroupMemberHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/add-group-member", http1.NewServer(endpoints.AddGroupMemberEndpoint, decodeAddGroupMemberRequest, encodeAddGroupMemberResponse, options...))
}

// decodeAddGroupMemberRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAddGroupMemberRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.AddGroupMemberRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeAddGroupMemberResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAddGroupMemberResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateGroupHandler creates the handler logic
func makeUpdateGroupHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-group", http1.NewServer(endpoints.UpdateGroupEndpoint, decodeUpdateGroupRequest, encodeUpdateGroupResponse, options...))
}

// decodeUpdateGroupRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateGroupRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateGroupResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateGroupResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateGroupMemberHandler creates the handler logic
func makeUpdateGroupMemberHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-group-member", http1.NewServer(endpoints.UpdateGroupMemberEndpoint, decodeUpdateGroupMemberRequest, encodeUpdateGroupMemberResponse, options...))
}

// decodeUpdateGroupMemberRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateGroupMemberRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateGroupMemberRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateGroupMemberResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateGroupMemberResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteGroupHandler creates the handler logic
func makeDeleteGroupHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-group", http1.NewServer(endpoints.DeleteGroupEndpoint, decodeDeleteGroupRequest, encodeDeleteGroupResponse, options...))
}

// decodeDeleteGroupRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteGroupRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteGroupResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteGroupResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteGroupMemberHandler creates the handler logic
func makeDeleteGroupMemberHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/delete-group-member", http1.NewServer(endpoints.DeleteGroupMemberEndpoint, decodeDeleteGroupMemberRequest, encodeDeleteGroupMemberResponse, options...))
}

// decodeDeleteGroupMemberRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteGroupMemberRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DeleteGroupMemberRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteGroupMemberResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteGroupMemberResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetGroupHandler creates the handler logic
func makeGetGroupHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-group", http1.NewServer(endpoints.GetGroupEndpoint, decodeGetGroupRequest, encodeGetGroupResponse, options...))
}

// decodeGetGroupRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetGroupRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetGroupResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetGroupResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetGroupsHandler creates the handler logic
func makeGetGroupsHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-groups", http1.NewServer(endpoints.GetGroupsEndpoint, decodeGetGroupsRequest, encodeGetGroupsResponse, options...))
}

// decodeGetGroupsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetGroupsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetGroupsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetGroupsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetGroupsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetGroupsByTypeHandler creates the handler logic
func makeGetGroupsByTypeHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-groups-by-type", http1.NewServer(endpoints.GetGroupsByTypeEndpoint, decodeGetGroupsByTypeRequest, encodeGetGroupsByTypeResponse, options...))
}

// decodeGetGroupsByTypeRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetGroupsByTypeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetGroupsByTypeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetGroupsByTypeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetGroupsByTypeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetGroupsByNumMembersHandler creates the handler logic
func makeGetGroupsByNumMembersHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-groups-by-num-members", http1.NewServer(endpoints.GetGroupsByNumMembersEndpoint, decodeGetGroupsByNumMembersRequest, encodeGetGroupsByNumMembersResponse, options...))
}

// decodeGetGroupsByNumMembersRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetGroupsByNumMembersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetGroupsByNumMembersRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetGroupsByNumMembersResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetGroupsByNumMembersResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetGroupsByTagsHandler creates the handler logic
func makeGetGroupsByTagsHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-groups-by-tags", http1.NewServer(endpoints.GetGroupsByTagsEndpoint, decodeGetGroupsByTagsRequest, encodeGetGroupsByTagsResponse, options...))
}

// decodeGetGroupsByTagsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetGroupsByTagsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetGroupsByTagsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetGroupsByTagsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetGroupsByTagsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
