package http

import (
	"bytes"
	"context"
	"encoding/json"
	endpoint1 "github.com/LensPlatform/micro/user-microservice/pkg/endpoint"
	http2 "github.com/LensPlatform/micro/user-microservice/pkg/http"
	service "github.com/LensPlatform/micro/user-microservice/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"
	"io/ioutil"
	http1 "net/http"
	"net/url"
	"strings"
)

// New returns an AddService backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".
func New(instance string, options map[string][]http.ClientOption) (service.MicroService, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	var createUserEndpoint endpoint.Endpoint
	{
		createUserEndpoint = http.NewClient("POST", copyURL(u, "/create-user"), encodeHTTPGenericRequest, decodeCreateUserResponse, options["CreateUser"]...).Endpoint()
	}

	var createProfileEndpoint endpoint.Endpoint
	{
		createProfileEndpoint = http.NewClient("POST", copyURL(u, "/create-profile"), encodeHTTPGenericRequest, decodeCreateProfileResponse, options["CreateProfile"]...).Endpoint()
	}

	var createSubscriptionEndpoint endpoint.Endpoint
	{
		createSubscriptionEndpoint = http.NewClient("POST", copyURL(u, "/create-subscription"), encodeHTTPGenericRequest, decodeCreateSubscriptionResponse, options["CreateSubscription"]...).Endpoint()
	}

	var createSubscriptionsEndpoint endpoint.Endpoint
	{
		createSubscriptionsEndpoint = http.NewClient("POST", copyURL(u, "/create-subscriptions"), encodeHTTPGenericRequest, decodeCreateSubscriptionsResponse, options["CreateSubscriptions"]...).Endpoint()
	}

	var updateUserEndpoint endpoint.Endpoint
	{
		updateUserEndpoint = http.NewClient("POST", copyURL(u, "/update-user"), encodeHTTPGenericRequest, decodeUpdateUserResponse, options["UpdateUser"]...).Endpoint()
	}

	var updateUserProfileEndpoint endpoint.Endpoint
	{
		updateUserProfileEndpoint = http.NewClient("POST", copyURL(u, "/update-user-profile"), encodeHTTPGenericRequest, decodeUpdateUserProfileResponse, options["UpdateUserProfile"]...).Endpoint()
	}

	var updateUserSubscriptionEndpoint endpoint.Endpoint
	{
		updateUserSubscriptionEndpoint = http.NewClient("POST", copyURL(u, "/update-user-subscription"), encodeHTTPGenericRequest, decodeUpdateUserSubscriptionResponse, options["UpdateUserSubscription"]...).Endpoint()
	}

	var deleteUserEndpoint endpoint.Endpoint
	{
		deleteUserEndpoint = http.NewClient("POST", copyURL(u, "/delete-user"), encodeHTTPGenericRequest, decodeDeleteUserResponse, options["DeleteUser"]...).Endpoint()
	}

	var deleteUserProfileEndpoint endpoint.Endpoint
	{
		deleteUserProfileEndpoint = http.NewClient("POST", copyURL(u, "/delete-user-profile"), encodeHTTPGenericRequest, decodeDeleteUserProfileResponse, options["DeleteUserProfile"]...).Endpoint()
	}

	var deleteSubscriptionEndpoint endpoint.Endpoint
	{
		deleteSubscriptionEndpoint = http.NewClient("POST", copyURL(u, "/delete-subscription"), encodeHTTPGenericRequest, decodeDeleteSubscriptionResponse, options["DeleteSubscription"]...).Endpoint()
	}

	var getUserEndpoint endpoint.Endpoint
	{
		getUserEndpoint = http.NewClient("POST", copyURL(u, "/get-user"), encodeHTTPGenericRequest, decodeGetUserResponse, options["GetUser"]...).Endpoint()
	}

	var getUsersEndpoint endpoint.Endpoint
	{
		getUsersEndpoint = http.NewClient("POST", copyURL(u, "/get-users"), encodeHTTPGenericRequest, decodeGetUsersResponse, options["GetUsers"]...).Endpoint()
	}

	var getUsersByAccountTypeEndpoint endpoint.Endpoint
	{
		getUsersByAccountTypeEndpoint = http.NewClient("POST", copyURL(u, "/get-users-by-account-type"), encodeHTTPGenericRequest, decodeGetUsersByAccountTypeResponse, options["GetUsersByAccountType"]...).Endpoint()
	}

	var getUsersByIntentEndpoint endpoint.Endpoint
	{
		getUsersByIntentEndpoint = http.NewClient("POST", copyURL(u, "/get-users-by-intent"), encodeHTTPGenericRequest, decodeGetUsersByIntentResponse, options["GetUsersByIntent"]...).Endpoint()
	}

	var getUserProfileEndpoint endpoint.Endpoint
	{
		getUserProfileEndpoint = http.NewClient("POST", copyURL(u, "/get-user-profile"), encodeHTTPGenericRequest, decodeGetUserProfileResponse, options["GetUserProfile"]...).Endpoint()
	}

	var getUserProfilesEndpoint endpoint.Endpoint
	{
		getUserProfilesEndpoint = http.NewClient("POST", copyURL(u, "/get-user-profiles"), encodeHTTPGenericRequest, decodeGetUserProfilesResponse, options["GetUserProfiles"]...).Endpoint()
	}

	var getUserProfilesByTypeEndpoint endpoint.Endpoint
	{
		getUserProfilesByTypeEndpoint = http.NewClient("POST", copyURL(u, "/get-user-profiles-by-type"), encodeHTTPGenericRequest, decodeGetUserProfilesByTypeResponse, options["GetUserProfilesByType"]...).Endpoint()
	}

	var getUserProfilesByNationalityEndpoint endpoint.Endpoint
	{
		getUserProfilesByNationalityEndpoint = http.NewClient("POST", copyURL(u, "/get-user-profiles-by-nationality"), encodeHTTPGenericRequest, decodeGetUserProfilesByNationalityResponse, options["GetUserProfilesByNationality"]...).Endpoint()
	}

	var getUserSubscriptionsEndpoint endpoint.Endpoint
	{
		getUserSubscriptionsEndpoint = http.NewClient("POST", copyURL(u, "/get-user-subscriptions"), encodeHTTPGenericRequest, decodeGetUserSubscriptionsResponse, options["GetUserSubscriptions"]...).Endpoint()
	}

	var createTeamEndpoint endpoint.Endpoint
	{
		createTeamEndpoint = http.NewClient("POST", copyURL(u, "/create-team"), encodeHTTPGenericRequest, decodeCreateTeamResponse, options["CreateTeam"]...).Endpoint()
	}

	var creatTeamProfileEndpoint endpoint.Endpoint
	{
		creatTeamProfileEndpoint = http.NewClient("POST", copyURL(u, "/creat-team-profile"), encodeHTTPGenericRequest, decodeCreatTeamProfileResponse, options["CreatTeamProfile"]...).Endpoint()
	}

	var createTeamSubscriptionEndpoint endpoint.Endpoint
	{
		createTeamSubscriptionEndpoint = http.NewClient("POST", copyURL(u, "/create-team-subscription"), encodeHTTPGenericRequest, decodeCreateTeamSubscriptionResponse, options["CreateTeamSubscription"]...).Endpoint()
	}

	var addMemberToTeamEndpoint endpoint.Endpoint
	{
		addMemberToTeamEndpoint = http.NewClient("POST", copyURL(u, "/add-member-to-team"), encodeHTTPGenericRequest, decodeAddMemberToTeamResponse, options["AddMemberToTeam"]...).Endpoint()
	}

	var addAdvisorToTeamEndpoint endpoint.Endpoint
	{
		addAdvisorToTeamEndpoint = http.NewClient("POST", copyURL(u, "/add-advisor-to-team"), encodeHTTPGenericRequest, decodeAddAdvisorToTeamResponse, options["AddAdvisorToTeam"]...).Endpoint()
	}

	var updateTeamEndpoint endpoint.Endpoint
	{
		updateTeamEndpoint = http.NewClient("POST", copyURL(u, "/update-team"), encodeHTTPGenericRequest, decodeUpdateTeamResponse, options["UpdateTeam"]...).Endpoint()
	}

	var updateTeamAdminEndpoint endpoint.Endpoint
	{
		updateTeamAdminEndpoint = http.NewClient("POST", copyURL(u, "/update-team-admin"), encodeHTTPGenericRequest, decodeUpdateTeamAdminResponse, options["UpdateTeamAdmin"]...).Endpoint()
	}

	var updateTeamMemberEndpoint endpoint.Endpoint
	{
		updateTeamMemberEndpoint = http.NewClient("POST", copyURL(u, "/update-team-member"), encodeHTTPGenericRequest, decodeUpdateTeamMemberResponse, options["UpdateTeamMember"]...).Endpoint()
	}

	var updateTeamAdvisorEndpoint endpoint.Endpoint
	{
		updateTeamAdvisorEndpoint = http.NewClient("POST", copyURL(u, "/update-team-advisor"), encodeHTTPGenericRequest, decodeUpdateTeamAdvisorResponse, options["UpdateTeamAdvisor"]...).Endpoint()
	}

	var updateTeamProfileEndpoint endpoint.Endpoint
	{
		updateTeamProfileEndpoint = http.NewClient("POST", copyURL(u, "/update-team-profile"), encodeHTTPGenericRequest, decodeUpdateTeamProfileResponse, options["UpdateTeamProfile"]...).Endpoint()
	}

	var updateTeamSubscriptionEndpoint endpoint.Endpoint
	{
		updateTeamSubscriptionEndpoint = http.NewClient("POST", copyURL(u, "/update-team-subscription"), encodeHTTPGenericRequest, decodeUpdateTeamSubscriptionResponse, options["UpdateTeamSubscription"]...).Endpoint()
	}

	var deleteTeamEndpoint endpoint.Endpoint
	{
		deleteTeamEndpoint = http.NewClient("POST", copyURL(u, "/delete-team"), encodeHTTPGenericRequest, decodeDeleteTeamResponse, options["DeleteTeam"]...).Endpoint()
	}

	var deleteTeamProfileEndpoint endpoint.Endpoint
	{
		deleteTeamProfileEndpoint = http.NewClient("POST", copyURL(u, "/delete-team-profile"), encodeHTTPGenericRequest, decodeDeleteTeamProfileResponse, options["DeleteTeamProfile"]...).Endpoint()
	}

	var deleteTeamMemberEndpoint endpoint.Endpoint
	{
		deleteTeamMemberEndpoint = http.NewClient("POST", copyURL(u, "/delete-team-member"), encodeHTTPGenericRequest, decodeDeleteTeamMemberResponse, options["DeleteTeamMember"]...).Endpoint()
	}

	var deleteTeamAdvisorEndpoint endpoint.Endpoint
	{
		deleteTeamAdvisorEndpoint = http.NewClient("POST", copyURL(u, "/delete-team-advisor"), encodeHTTPGenericRequest, decodeDeleteTeamAdvisorResponse, options["DeleteTeamAdvisor"]...).Endpoint()
	}

	var deleteTeamAdminEndpoint endpoint.Endpoint
	{
		deleteTeamAdminEndpoint = http.NewClient("POST", copyURL(u, "/delete-team-admin"), encodeHTTPGenericRequest, decodeDeleteTeamAdminResponse, options["DeleteTeamAdmin"]...).Endpoint()
	}

	var deleteTeamSubscriptionEndpoint endpoint.Endpoint
	{
		deleteTeamSubscriptionEndpoint = http.NewClient("POST", copyURL(u, "/delete-team-subscription"), encodeHTTPGenericRequest, decodeDeleteTeamSubscriptionResponse, options["DeleteTeamSubscription"]...).Endpoint()
	}

	var getTeamEndpoint endpoint.Endpoint
	{
		getTeamEndpoint = http.NewClient("POST", copyURL(u, "/get-team"), encodeHTTPGenericRequest, decodeGetTeamResponse, options["GetTeam"]...).Endpoint()
	}

	var getTeamProfileEndpoint endpoint.Endpoint
	{
		getTeamProfileEndpoint = http.NewClient("POST", copyURL(u, "/get-team-profile"), encodeHTTPGenericRequest, decodeGetTeamProfileResponse, options["GetTeamProfile"]...).Endpoint()
	}

	var getTeamSubscriptionsEndpoint endpoint.Endpoint
	{
		getTeamSubscriptionsEndpoint = http.NewClient("POST", copyURL(u, "/get-team-subscriptions"), encodeHTTPGenericRequest, decodeGetTeamSubscriptionsResponse, options["GetTeamSubscriptions"]...).Endpoint()
	}

	var getTeamMembersEndpoint endpoint.Endpoint
	{
		getTeamMembersEndpoint = http.NewClient("POST", copyURL(u, "/get-team-members"), encodeHTTPGenericRequest, decodeGetTeamMembersResponse, options["GetTeamMembers"]...).Endpoint()
	}

	var getTeamAdvisorsEndpoint endpoint.Endpoint
	{
		getTeamAdvisorsEndpoint = http.NewClient("POST", copyURL(u, "/get-team-advisors"), encodeHTTPGenericRequest, decodeGetTeamAdvisorsResponse, options["GetTeamAdvisors"]...).Endpoint()
	}

	var getTeamAdminEndpoint endpoint.Endpoint
	{
		getTeamAdminEndpoint = http.NewClient("POST", copyURL(u, "/get-team-admin"), encodeHTTPGenericRequest, decodeGetTeamAdminResponse, options["GetTeamAdmin"]...).Endpoint()
	}

	var getTeamsEndpoint endpoint.Endpoint
	{
		getTeamsEndpoint = http.NewClient("POST", copyURL(u, "/get-teams"), encodeHTTPGenericRequest, decodeGetTeamsResponse, options["GetTeams"]...).Endpoint()
	}

	var getTeamsByTypeEndpoint endpoint.Endpoint
	{
		getTeamsByTypeEndpoint = http.NewClient("POST", copyURL(u, "/get-teams-by-type"), encodeHTTPGenericRequest, decodeGetTeamsByTypeResponse, options["GetTeamsByType"]...).Endpoint()
	}

	var getTeamsByIndustryEndpoint endpoint.Endpoint
	{
		getTeamsByIndustryEndpoint = http.NewClient("POST", copyURL(u, "/get-teams-by-industry"), encodeHTTPGenericRequest, decodeGetTeamsByIndustryResponse, options["GetTeamsByIndustry"]...).Endpoint()
	}

	var getTeamsByNumberOfEmployeesEndpoint endpoint.Endpoint
	{
		getTeamsByNumberOfEmployeesEndpoint = http.NewClient("POST", copyURL(u, "/get-teams-by-number-of-employees"), encodeHTTPGenericRequest, decodeGetTeamsByNumberOfEmployeesResponse, options["GetTeamsByNumberOfEmployees"]...).Endpoint()
	}

	var getTeamProfilesEndpoint endpoint.Endpoint
	{
		getTeamProfilesEndpoint = http.NewClient("POST", copyURL(u, "/get-team-profiles"), encodeHTTPGenericRequest, decodeGetTeamProfilesResponse, options["GetTeamProfiles"]...).Endpoint()
	}

	var getTeamsByTagsEndpoint endpoint.Endpoint
	{
		getTeamsByTagsEndpoint = http.NewClient("POST", copyURL(u, "/get-teams-by-tags"), encodeHTTPGenericRequest, decodeGetTeamsByTagsResponse, options["GetTeamsByTags"]...).Endpoint()
	}

	var createGroupEndpoint endpoint.Endpoint
	{
		createGroupEndpoint = http.NewClient("POST", copyURL(u, "/create-group"), encodeHTTPGenericRequest, decodeCreateGroupResponse, options["CreateGroup"]...).Endpoint()
	}

	var addGroupMemberEndpoint endpoint.Endpoint
	{
		addGroupMemberEndpoint = http.NewClient("POST", copyURL(u, "/add-group-member"), encodeHTTPGenericRequest, decodeAddGroupMemberResponse, options["AddGroupMember"]...).Endpoint()
	}

	var updateGroupEndpoint endpoint.Endpoint
	{
		updateGroupEndpoint = http.NewClient("POST", copyURL(u, "/update-group"), encodeHTTPGenericRequest, decodeUpdateGroupResponse, options["UpdateGroup"]...).Endpoint()
	}

	var updateGroupMemberEndpoint endpoint.Endpoint
	{
		updateGroupMemberEndpoint = http.NewClient("POST", copyURL(u, "/update-group-member"), encodeHTTPGenericRequest, decodeUpdateGroupMemberResponse, options["UpdateGroupMember"]...).Endpoint()
	}

	var deleteGroupEndpoint endpoint.Endpoint
	{
		deleteGroupEndpoint = http.NewClient("POST", copyURL(u, "/delete-group"), encodeHTTPGenericRequest, decodeDeleteGroupResponse, options["DeleteGroup"]...).Endpoint()
	}

	var deleteGroupMemberEndpoint endpoint.Endpoint
	{
		deleteGroupMemberEndpoint = http.NewClient("POST", copyURL(u, "/delete-group-member"), encodeHTTPGenericRequest, decodeDeleteGroupMemberResponse, options["DeleteGroupMember"]...).Endpoint()
	}

	var getGroupEndpoint endpoint.Endpoint
	{
		getGroupEndpoint = http.NewClient("POST", copyURL(u, "/get-group"), encodeHTTPGenericRequest, decodeGetGroupResponse, options["GetGroup"]...).Endpoint()
	}

	var getGroupsEndpoint endpoint.Endpoint
	{
		getGroupsEndpoint = http.NewClient("POST", copyURL(u, "/get-groups"), encodeHTTPGenericRequest, decodeGetGroupsResponse, options["GetGroups"]...).Endpoint()
	}

	var getGroupsByTypeEndpoint endpoint.Endpoint
	{
		getGroupsByTypeEndpoint = http.NewClient("POST", copyURL(u, "/get-groups-by-type"), encodeHTTPGenericRequest, decodeGetGroupsByTypeResponse, options["GetGroupsByType"]...).Endpoint()
	}

	var getGroupsByNumMembersEndpoint endpoint.Endpoint
	{
		getGroupsByNumMembersEndpoint = http.NewClient("POST", copyURL(u, "/get-groups-by-num-members"), encodeHTTPGenericRequest, decodeGetGroupsByNumMembersResponse, options["GetGroupsByNumMembers"]...).Endpoint()
	}

	var getGroupsByTagsEndpoint endpoint.Endpoint
	{
		getGroupsByTagsEndpoint = http.NewClient("POST", copyURL(u, "/get-groups-by-tags"), encodeHTTPGenericRequest, decodeGetGroupsByTagsResponse, options["GetGroupsByTags"]...).Endpoint()
	}

	return endpoint1.Endpoints{
		AddAdvisorToTeamEndpoint:             addAdvisorToTeamEndpoint,
		AddGroupMemberEndpoint:               addGroupMemberEndpoint,
		AddMemberToTeamEndpoint:              addMemberToTeamEndpoint,
		CreatTeamProfileEndpoint:             creatTeamProfileEndpoint,
		CreateGroupEndpoint:                  createGroupEndpoint,
		CreateProfileEndpoint:                createProfileEndpoint,
		CreateSubscriptionEndpoint:           createSubscriptionEndpoint,
		CreateSubscriptionsEndpoint:          createSubscriptionsEndpoint,
		CreateTeamEndpoint:                   createTeamEndpoint,
		CreateTeamSubscriptionEndpoint:       createTeamSubscriptionEndpoint,
		CreateUserEndpoint:                   createUserEndpoint,
		DeleteGroupEndpoint:                  deleteGroupEndpoint,
		DeleteGroupMemberEndpoint:            deleteGroupMemberEndpoint,
		DeleteSubscriptionEndpoint:           deleteSubscriptionEndpoint,
		DeleteTeamAdminEndpoint:              deleteTeamAdminEndpoint,
		DeleteTeamAdvisorEndpoint:            deleteTeamAdvisorEndpoint,
		DeleteTeamEndpoint:                   deleteTeamEndpoint,
		DeleteTeamMemberEndpoint:             deleteTeamMemberEndpoint,
		DeleteTeamProfileEndpoint:            deleteTeamProfileEndpoint,
		DeleteTeamSubscriptionEndpoint:       deleteTeamSubscriptionEndpoint,
		DeleteUserEndpoint:                   deleteUserEndpoint,
		DeleteUserProfileEndpoint:            deleteUserProfileEndpoint,
		GetGroupEndpoint:                     getGroupEndpoint,
		GetGroupsByNumMembersEndpoint:        getGroupsByNumMembersEndpoint,
		GetGroupsByTagsEndpoint:              getGroupsByTagsEndpoint,
		GetGroupsByTypeEndpoint:              getGroupsByTypeEndpoint,
		GetGroupsEndpoint:                    getGroupsEndpoint,
		GetTeamAdminEndpoint:                 getTeamAdminEndpoint,
		GetTeamAdvisorsEndpoint:              getTeamAdvisorsEndpoint,
		GetTeamEndpoint:                      getTeamEndpoint,
		GetTeamMembersEndpoint:               getTeamMembersEndpoint,
		GetTeamProfileEndpoint:               getTeamProfileEndpoint,
		GetTeamProfilesEndpoint:              getTeamProfilesEndpoint,
		GetTeamSubscriptionsEndpoint:         getTeamSubscriptionsEndpoint,
		GetTeamsByIndustryEndpoint:           getTeamsByIndustryEndpoint,
		GetTeamsByNumberOfEmployeesEndpoint:  getTeamsByNumberOfEmployeesEndpoint,
		GetTeamsByTagsEndpoint:               getTeamsByTagsEndpoint,
		GetTeamsByTypeEndpoint:               getTeamsByTypeEndpoint,
		GetTeamsEndpoint:                     getTeamsEndpoint,
		GetUserEndpoint:                      getUserEndpoint,
		GetUserProfileEndpoint:               getUserProfileEndpoint,
		GetUserProfilesByNationalityEndpoint: getUserProfilesByNationalityEndpoint,
		GetUserProfilesByTypeEndpoint:        getUserProfilesByTypeEndpoint,
		GetUserProfilesEndpoint:              getUserProfilesEndpoint,
		GetUserSubscriptionsEndpoint:         getUserSubscriptionsEndpoint,
		GetUsersByAccountTypeEndpoint:        getUsersByAccountTypeEndpoint,
		GetUsersByIntentEndpoint:             getUsersByIntentEndpoint,
		GetUsersEndpoint:                     getUsersEndpoint,
		UpdateGroupEndpoint:                  updateGroupEndpoint,
		UpdateGroupMemberEndpoint:            updateGroupMemberEndpoint,
		UpdateTeamAdminEndpoint:              updateTeamAdminEndpoint,
		UpdateTeamAdvisorEndpoint:            updateTeamAdvisorEndpoint,
		UpdateTeamEndpoint:                   updateTeamEndpoint,
		UpdateTeamMemberEndpoint:             updateTeamMemberEndpoint,
		UpdateTeamProfileEndpoint:            updateTeamProfileEndpoint,
		UpdateTeamSubscriptionEndpoint:       updateTeamSubscriptionEndpoint,
		UpdateUserEndpoint:                   updateUserEndpoint,
		UpdateUserProfileEndpoint:            updateUserProfileEndpoint,
		UpdateUserSubscriptionEndpoint:       updateUserSubscriptionEndpoint,
	}, nil
}

// EncodeHTTPGenericRequest is a transport/http.EncodeRequestFunc that
// SON-encodes any request to the request body. Primarily useful in a client.
func encodeHTTPGenericRequest(_ context.Context, r *http1.Request, request interface{}) error {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// decodeCreateUserResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeCreateUserResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.CreateUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeCreateProfileResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeCreateProfileResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.CreateProfileResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeCreateSubscriptionResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeCreateSubscriptionResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.CreateSubscriptionResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeCreateSubscriptionsResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeCreateSubscriptionsResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.CreateSubscriptionsResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateUserResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateUserResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateUserProfileResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateUserProfileResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateUserProfileResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateUserSubscriptionResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateUserSubscriptionResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateUserSubscriptionResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteUserResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteUserResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteUserProfileResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteUserProfileResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteUserProfileResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteSubscriptionResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteSubscriptionResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteSubscriptionResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetUserResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetUserResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetUserResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetUsersResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetUsersResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetUsersResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetUsersByAccountTypeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetUsersByAccountTypeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetUsersByAccountTypeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetUsersByIntentResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetUsersByIntentResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetUsersByIntentResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetUserProfileResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetUserProfileResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetUserProfileResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetUserProfilesResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetUserProfilesResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetUserProfilesResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetUserProfilesByTypeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetUserProfilesByTypeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetUserProfilesByTypeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetUserProfilesByNationalityResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetUserProfilesByNationalityResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetUserProfilesByNationalityResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetUserSubscriptionsResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetUserSubscriptionsResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetUserSubscriptionsResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeCreateTeamResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeCreateTeamResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.CreateTeamResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeCreatTeamProfileResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeCreatTeamProfileResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.CreatTeamProfileResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeCreateTeamSubscriptionResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeCreateTeamSubscriptionResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.CreateTeamSubscriptionResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeAddMemberToTeamResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeAddMemberToTeamResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.AddMemberToTeamResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeAddAdvisorToTeamResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeAddAdvisorToTeamResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.AddAdvisorToTeamResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateTeamResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateTeamResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateTeamResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateTeamAdminResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateTeamAdminResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateTeamAdminResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateTeamMemberResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateTeamMemberResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateTeamMemberResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateTeamAdvisorResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateTeamAdvisorResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateTeamAdvisorResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateTeamProfileResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateTeamProfileResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateTeamProfileResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateTeamSubscriptionResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateTeamSubscriptionResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateTeamSubscriptionResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteTeamResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteTeamResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteTeamResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteTeamProfileResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteTeamProfileResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteTeamProfileResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteTeamMemberResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteTeamMemberResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteTeamMemberResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteTeamAdvisorResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteTeamAdvisorResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteTeamAdvisorResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteTeamAdminResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteTeamAdminResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteTeamAdminResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteTeamSubscriptionResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteTeamSubscriptionResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteTeamSubscriptionResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetTeamResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetTeamResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetTeamResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetTeamProfileResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetTeamProfileResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetTeamProfileResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetTeamSubscriptionsResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetTeamSubscriptionsResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetTeamSubscriptionsResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetTeamMembersResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetTeamMembersResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetTeamMembersResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetTeamAdvisorsResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetTeamAdvisorsResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetTeamAdvisorsResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetTeamAdminResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetTeamAdminResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetTeamAdminResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetTeamsResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetTeamsResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetTeamsResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetTeamsByTypeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetTeamsByTypeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetTeamsByTypeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetTeamsByIndustryResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetTeamsByIndustryResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetTeamsByIndustryResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetTeamsByNumberOfEmployeesResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetTeamsByNumberOfEmployeesResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetTeamsByNumberOfEmployeesResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetTeamProfilesResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetTeamProfilesResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetTeamProfilesResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetTeamsByTagsResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetTeamsByTagsResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetTeamsByTagsResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeCreateGroupResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeCreateGroupResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.CreateGroupResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeAddGroupMemberResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeAddGroupMemberResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.AddGroupMemberResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateGroupResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateGroupResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateGroupResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateGroupMemberResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateGroupMemberResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateGroupMemberResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteGroupResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteGroupResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteGroupResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteGroupMemberResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteGroupMemberResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteGroupMemberResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetGroupResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetGroupResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetGroupResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetGroupsResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetGroupsResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetGroupsResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetGroupsByTypeResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetGroupsByTypeResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetGroupsByTypeResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetGroupsByNumMembersResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetGroupsByNumMembersResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetGroupsByNumMembersResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetGroupsByTagsResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetGroupsByTagsResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetGroupsByTagsResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
func copyURL(base *url.URL, path string) (next *url.URL) {
	n := *base
	n.Path = path
	next = &n
	return
}
