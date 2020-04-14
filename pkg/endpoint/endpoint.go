package endpoint

import (
	"context"

	model "github.com/LensPlatform/micro/pkg/models/proto"
	proto "github.com/LensPlatform/micro/pkg/models/proto"
	service "github.com/LensPlatform/micro/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateUserRequest collects the request parameters for the CreateUser method.
type CreateUserRequest struct {
	User proto.User `json:"user"`
}

// CreateUserResponse collects the response parameters for the CreateUser method.
type CreateUserResponse struct {
	E0 error `json:"e0"`
}

// MakeCreateUserEndpoint returns an endpoint that invokes CreateUser on the service.
func MakeCreateUserEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		e0 := s.CreateUser(ctx, req.User)
		return CreateUserResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r CreateUserResponse) Failed() error {
	return r.E0
}

// CreateProfileRequest collects the request parameters for the CreateProfile method.
type CreateProfileRequest struct {
	UserId  int32         `json:"user_id"`
	Profile proto.Profile `json:"profile"`
}

// CreateProfileResponse collects the response parameters for the CreateProfile method.
type CreateProfileResponse struct {
	E0 error `json:"e0"`
}

// MakeCreateProfileEndpoint returns an endpoint that invokes CreateProfile on the service.
func MakeCreateProfileEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateProfileRequest)
		e0 := s.CreateProfile(ctx, req.UserId, req.Profile)
		return CreateProfileResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r CreateProfileResponse) Failed() error {
	return r.E0
}

// CreateSubscriptionRequest collects the request parameters for the CreateSubscription method.
type CreateSubscriptionRequest struct {
	UserId       int32               `json:"user_id"`
	Subscription proto.Subscriptions `json:"subscription"`
}

// CreateSubscriptionResponse collects the response parameters for the CreateSubscription method.
type CreateSubscriptionResponse struct {
	E0 error `json:"e0"`
}

// MakeCreateSubscriptionEndpoint returns an endpoint that invokes CreateSubscription on the service.
func MakeCreateSubscriptionEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateSubscriptionRequest)
		e0 := s.CreateSubscription(ctx, req.UserId, req.Subscription)
		return CreateSubscriptionResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r CreateSubscriptionResponse) Failed() error {
	return r.E0
}

// CreateSubscriptionsRequest collects the request parameters for the CreateSubscriptions method.
type CreateSubscriptionsRequest struct {
	UserId        []model.Subscriptions `json:"user_id"`
	Subscriptions []model.Subscriptions `json:"subscriptions"`
}

// CreateSubscriptionsResponse collects the response parameters for the CreateSubscriptions method.
type CreateSubscriptionsResponse struct {
	E0 error `json:"e0"`
}

// MakeCreateSubscriptionsEndpoint returns an endpoint that invokes CreateSubscriptions on the service.
func MakeCreateSubscriptionsEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateSubscriptionsRequest)
		e0 := s.CreateSubscriptions(ctx, req.UserId, req.Subscriptions)
		return CreateSubscriptionsResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r CreateSubscriptionsResponse) Failed() error {
	return r.E0
}

// UpdateUserRequest collects the request parameters for the UpdateUser method.
type UpdateUserRequest struct {
	UserId int32      `json:"user_id"`
	User   proto.User `json:"user"`
}

// UpdateUserResponse collects the response parameters for the UpdateUser method.
type UpdateUserResponse struct {
	E0 error      `json:"e0"`
	M1 proto.User `json:"m1"`
}

// MakeUpdateUserEndpoint returns an endpoint that invokes UpdateUser on the service.
func MakeUpdateUserEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		e0, m1 := s.UpdateUser(ctx, req.UserId, req.User)
		return UpdateUserResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateUserResponse) Failed() error {
	return r.E0
}

// UpdateUserProfileRequest collects the request parameters for the UpdateUserProfile method.
type UpdateUserProfileRequest struct {
	UserId    int32         `json:"user_id"`
	ProfileId int32         `json:"profile_id"`
	Profile   proto.Profile `json:"profile"`
}

// UpdateUserProfileResponse collects the response parameters for the UpdateUserProfile method.
type UpdateUserProfileResponse struct {
	E0 error         `json:"e0"`
	M1 proto.Profile `json:"m1"`
}

// MakeUpdateUserProfileEndpoint returns an endpoint that invokes UpdateUserProfile on the service.
func MakeUpdateUserProfileEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserProfileRequest)
		e0, m1 := s.UpdateUserProfile(ctx, req.UserId, req.ProfileId, req.Profile)
		return UpdateUserProfileResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateUserProfileResponse) Failed() error {
	return r.E0
}

// UpdateUserSubscriptionRequest collects the request parameters for the UpdateUserSubscription method.
type UpdateUserSubscriptionRequest struct {
	UserId         int32               `json:"user_id"`
	SubscriptionId int32               `json:"subscription_id"`
	Subscription   proto.Subscriptions `json:"subscription"`
}

// UpdateUserSubscriptionResponse collects the response parameters for the UpdateUserSubscription method.
type UpdateUserSubscriptionResponse struct {
	E0 error               `json:"e0"`
	M1 proto.Subscriptions `json:"m1"`
}

// MakeUpdateUserSubscriptionEndpoint returns an endpoint that invokes UpdateUserSubscription on the service.
func MakeUpdateUserSubscriptionEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserSubscriptionRequest)
		e0, m1 := s.UpdateUserSubscription(ctx, req.UserId, req.SubscriptionId, req.Subscription)
		return UpdateUserSubscriptionResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateUserSubscriptionResponse) Failed() error {
	return r.E0
}

// DeleteUserRequest collects the request parameters for the DeleteUser method.
type DeleteUserRequest struct {
	UserId int32 `json:"user_id"`
}

// DeleteUserResponse collects the response parameters for the DeleteUser method.
type DeleteUserResponse struct {
	E0 error `json:"e0"`
}

// MakeDeleteUserEndpoint returns an endpoint that invokes DeleteUser on the service.
func MakeDeleteUserEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)
		e0 := s.DeleteUser(ctx, req.UserId)
		return DeleteUserResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r DeleteUserResponse) Failed() error {
	return r.E0
}

// DeleteUserProfileRequest collects the request parameters for the DeleteUserProfile method.
type DeleteUserProfileRequest struct {
	UserId    int32 `json:"user_id"`
	ProfileId int32 `json:"profile_id"`
}

// DeleteUserProfileResponse collects the response parameters for the DeleteUserProfile method.
type DeleteUserProfileResponse struct {
	E0 error `json:"e0"`
}

// MakeDeleteUserProfileEndpoint returns an endpoint that invokes DeleteUserProfile on the service.
func MakeDeleteUserProfileEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserProfileRequest)
		e0 := s.DeleteUserProfile(ctx, req.UserId, req.ProfileId)
		return DeleteUserProfileResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r DeleteUserProfileResponse) Failed() error {
	return r.E0
}

// DeleteSubscriptionRequest collects the request parameters for the DeleteSubscription method.
type DeleteSubscriptionRequest struct {
	UserId         int32 `json:"user_id"`
	SubscriptionId int32 `json:"subscription_id"`
}

// DeleteSubscriptionResponse collects the response parameters for the DeleteSubscription method.
type DeleteSubscriptionResponse struct {
	E0 error `json:"e0"`
}

// MakeDeleteSubscriptionEndpoint returns an endpoint that invokes DeleteSubscription on the service.
func MakeDeleteSubscriptionEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteSubscriptionRequest)
		e0 := s.DeleteSubscription(ctx, req.UserId, req.SubscriptionId)
		return DeleteSubscriptionResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r DeleteSubscriptionResponse) Failed() error {
	return r.E0
}

// GetUserRequest collects the request parameters for the GetUser method.
type GetUserRequest struct {
	UserId int32 `json:"user_id"`
}

// GetUserResponse collects the response parameters for the GetUser method.
type GetUserResponse struct {
	E0 error      `json:"e0"`
	M1 proto.User `json:"m1"`
}

// MakeGetUserEndpoint returns an endpoint that invokes GetUser on the service.
func MakeGetUserEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		e0, m1 := s.GetUser(ctx, req.UserId)
		return GetUserResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserResponse) Failed() error {
	return r.E0
}

// GetUsersRequest collects the request parameters for the GetUsers method.
type GetUsersRequest struct {
	Limit int32 `json:"limit"`
}

// GetUsersResponse collects the response parameters for the GetUsers method.
type GetUsersResponse struct {
	E0 error        `json:"e0"`
	M1 []model.User `json:"m1"`
}

// MakeGetUsersEndpoint returns an endpoint that invokes GetUsers on the service.
func MakeGetUsersEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUsersRequest)
		e0, m1 := s.GetUsers(ctx, req.Limit)
		return GetUsersResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUsersResponse) Failed() error {
	return r.E0
}

// GetUsersByAccountTypeRequest collects the request parameters for the GetUsersByAccountType method.
type GetUsersByAccountTypeRequest struct {
	AccounType string `json:"accoun_type"`
	Limit      int32  `json:"limit"`
}

// GetUsersByAccountTypeResponse collects the response parameters for the GetUsersByAccountType method.
type GetUsersByAccountTypeResponse struct {
	E0 error        `json:"e0"`
	M1 []model.User `json:"m1"`
}

// MakeGetUsersByAccountTypeEndpoint returns an endpoint that invokes GetUsersByAccountType on the service.
func MakeGetUsersByAccountTypeEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUsersByAccountTypeRequest)
		e0, m1 := s.GetUsersByAccountType(ctx, req.AccounType, req.Limit)
		return GetUsersByAccountTypeResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUsersByAccountTypeResponse) Failed() error {
	return r.E0
}

// GetUsersByIntentRequest collects the request parameters for the GetUsersByIntent method.
type GetUsersByIntentRequest struct {
	Intent string `json:"intent"`
	Limit  int32  `json:"limit"`
}

// GetUsersByIntentResponse collects the response parameters for the GetUsersByIntent method.
type GetUsersByIntentResponse struct {
	E0 error        `json:"e0"`
	M1 []model.User `json:"m1"`
}

// MakeGetUsersByIntentEndpoint returns an endpoint that invokes GetUsersByIntent on the service.
func MakeGetUsersByIntentEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUsersByIntentRequest)
		e0, m1 := s.GetUsersByIntent(ctx, req.Intent, req.Limit)
		return GetUsersByIntentResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUsersByIntentResponse) Failed() error {
	return r.E0
}

// GetUserProfileRequest collects the request parameters for the GetUserProfile method.
type GetUserProfileRequest struct {
	UserId int32 `json:"user_id"`
}

// GetUserProfileResponse collects the response parameters for the GetUserProfile method.
type GetUserProfileResponse struct {
	E0 error         `json:"e0"`
	M1 proto.Profile `json:"m1"`
}

// MakeGetUserProfileEndpoint returns an endpoint that invokes GetUserProfile on the service.
func MakeGetUserProfileEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserProfileRequest)
		e0, m1 := s.GetUserProfile(ctx, req.UserId)
		return GetUserProfileResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserProfileResponse) Failed() error {
	return r.E0
}

// GetUserProfilesRequest collects the request parameters for the GetUserProfiles method.
type GetUserProfilesRequest struct {
	Limit int32 `json:"limit"`
}

// GetUserProfilesResponse collects the response parameters for the GetUserProfiles method.
type GetUserProfilesResponse struct {
	E0 error           `json:"e0"`
	M1 []model.Profile `json:"m1"`
}

// MakeGetUserProfilesEndpoint returns an endpoint that invokes GetUserProfiles on the service.
func MakeGetUserProfilesEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserProfilesRequest)
		e0, m1 := s.GetUserProfiles(ctx, req.Limit)
		return GetUserProfilesResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserProfilesResponse) Failed() error {
	return r.E0
}

// GetUserProfilesByTypeRequest collects the request parameters for the GetUserProfilesByType method.
type GetUserProfilesByTypeRequest struct {
	ProfileType string `json:"profile_type"`
	Limit       int32  `json:"limit"`
}

// GetUserProfilesByTypeResponse collects the response parameters for the GetUserProfilesByType method.
type GetUserProfilesByTypeResponse struct {
	E0 error           `json:"e0"`
	M1 []model.Profile `json:"m1"`
}

// MakeGetUserProfilesByTypeEndpoint returns an endpoint that invokes GetUserProfilesByType on the service.
func MakeGetUserProfilesByTypeEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserProfilesByTypeRequest)
		e0, m1 := s.GetUserProfilesByType(ctx, req.ProfileType, req.Limit)
		return GetUserProfilesByTypeResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserProfilesByTypeResponse) Failed() error {
	return r.E0
}

// GetUserProfilesByNationalityRequest collects the request parameters for the GetUserProfilesByNationality method.
type GetUserProfilesByNationalityRequest struct {
	Nationality string `json:"nationality"`
	Limit       int32  `json:"limit"`
}

// GetUserProfilesByNationalityResponse collects the response parameters for the GetUserProfilesByNationality method.
type GetUserProfilesByNationalityResponse struct {
	E0 error           `json:"e0"`
	M1 []model.Profile `json:"m1"`
}

// MakeGetUserProfilesByNationalityEndpoint returns an endpoint that invokes GetUserProfilesByNationality on the service.
func MakeGetUserProfilesByNationalityEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserProfilesByNationalityRequest)
		e0, m1 := s.GetUserProfilesByNationality(ctx, req.Nationality, req.Limit)
		return GetUserProfilesByNationalityResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserProfilesByNationalityResponse) Failed() error {
	return r.E0
}

// GetUserSubscriptionsRequest collects the request parameters for the GetUserSubscriptions method.
type GetUserSubscriptionsRequest struct {
	UserId int32 `json:"user_id"`
}

// GetUserSubscriptionsResponse collects the response parameters for the GetUserSubscriptions method.
type GetUserSubscriptionsResponse struct {
	E0 error                 `json:"e0"`
	M1 []model.Subscriptions `json:"m1"`
}

// MakeGetUserSubscriptionsEndpoint returns an endpoint that invokes GetUserSubscriptions on the service.
func MakeGetUserSubscriptionsEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserSubscriptionsRequest)
		e0, m1 := s.GetUserSubscriptions(ctx, req.UserId)
		return GetUserSubscriptionsResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserSubscriptionsResponse) Failed() error {
	return r.E0
}

// CreateTeamRequest collects the request parameters for the CreateTeam method.
type CreateTeamRequest struct {
	Team  proto.Team `json:"team"`
	Admin proto.User `json:"admin"`
}

// CreateTeamResponse collects the response parameters for the CreateTeam method.
type CreateTeamResponse struct {
	E0 error `json:"e0"`
}

// MakeCreateTeamEndpoint returns an endpoint that invokes CreateTeam on the service.
func MakeCreateTeamEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTeamRequest)
		e0 := s.CreateTeam(ctx, req.Team, req.Admin)
		return CreateTeamResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r CreateTeamResponse) Failed() error {
	return r.E0
}

// CreatTeamProfileRequest collects the request parameters for the CreatTeamProfile method.
type CreatTeamProfileRequest struct {
	TeamId  int32             `json:"team_id"`
	Profile proto.TeamProfile `json:"profile"`
}

// CreatTeamProfileResponse collects the response parameters for the CreatTeamProfile method.
type CreatTeamProfileResponse struct {
	E0 error `json:"e0"`
}

// MakeCreatTeamProfileEndpoint returns an endpoint that invokes CreatTeamProfile on the service.
func MakeCreatTeamProfileEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreatTeamProfileRequest)
		e0 := s.CreatTeamProfile(ctx, req.TeamId, req.Profile)
		return CreatTeamProfileResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r CreatTeamProfileResponse) Failed() error {
	return r.E0
}

// CreateTeamSubscriptionRequest collects the request parameters for the CreateTeamSubscription method.
type CreateTeamSubscriptionRequest struct {
	TeamId       int32               `json:"team_id"`
	Subscription proto.Subscriptions `json:"subscription"`
}

// CreateTeamSubscriptionResponse collects the response parameters for the CreateTeamSubscription method.
type CreateTeamSubscriptionResponse struct {
	E0 error `json:"e0"`
}

// MakeCreateTeamSubscriptionEndpoint returns an endpoint that invokes CreateTeamSubscription on the service.
func MakeCreateTeamSubscriptionEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTeamSubscriptionRequest)
		e0 := s.CreateTeamSubscription(ctx, req.TeamId, req.Subscription)
		return CreateTeamSubscriptionResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r CreateTeamSubscriptionResponse) Failed() error {
	return r.E0
}

// AddMemberToTeamRequest collects the request parameters for the AddMemberToTeam method.
type AddMemberToTeamRequest struct {
	TeamId int32      `json:"team_id"`
	Member proto.User `json:"member"`
}

// AddMemberToTeamResponse collects the response parameters for the AddMemberToTeam method.
type AddMemberToTeamResponse struct {
	E0 error `json:"e0"`
}

// MakeAddMemberToTeamEndpoint returns an endpoint that invokes AddMemberToTeam on the service.
func MakeAddMemberToTeamEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddMemberToTeamRequest)
		e0 := s.AddMemberToTeam(ctx, req.TeamId, req.Member)
		return AddMemberToTeamResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r AddMemberToTeamResponse) Failed() error {
	return r.E0
}

// AddAdvisorToTeamRequest collects the request parameters for the AddAdvisorToTeam method.
type AddAdvisorToTeamRequest struct {
	TeamId    int32      `json:"team_id"`
	AdvisorId int32      `json:"advisor_id"`
	Advisor   proto.User `json:"advisor"`
}

// AddAdvisorToTeamResponse collects the response parameters for the AddAdvisorToTeam method.
type AddAdvisorToTeamResponse struct {
	E0 error `json:"e0"`
}

// MakeAddAdvisorToTeamEndpoint returns an endpoint that invokes AddAdvisorToTeam on the service.
func MakeAddAdvisorToTeamEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddAdvisorToTeamRequest)
		e0 := s.AddAdvisorToTeam(ctx, req.TeamId, req.AdvisorId, req.Advisor)
		return AddAdvisorToTeamResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r AddAdvisorToTeamResponse) Failed() error {
	return r.E0
}

// UpdateTeamRequest collects the request parameters for the UpdateTeam method.
type UpdateTeamRequest struct {
	TeamId int32      `json:"team_id"`
	Team   proto.Team `json:"team"`
}

// UpdateTeamResponse collects the response parameters for the UpdateTeam method.
type UpdateTeamResponse struct {
	E0 error      `json:"e0"`
	M1 proto.Team `json:"m1"`
}

// MakeUpdateTeamEndpoint returns an endpoint that invokes UpdateTeam on the service.
func MakeUpdateTeamEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateTeamRequest)
		e0, m1 := s.UpdateTeam(ctx, req.TeamId, req.Team)
		return UpdateTeamResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateTeamResponse) Failed() error {
	return r.E0
}

// UpdateTeamAdminRequest collects the request parameters for the UpdateTeamAdmin method.
type UpdateTeamAdminRequest struct {
	TeamId  int32      `json:"team_id"`
	AdminId int32      `json:"admin_id"`
	Admin   proto.User `json:"admin"`
}

// UpdateTeamAdminResponse collects the response parameters for the UpdateTeamAdmin method.
type UpdateTeamAdminResponse struct {
	E0 error      `json:"e0"`
	M1 proto.Team `json:"m1"`
}

// MakeUpdateTeamAdminEndpoint returns an endpoint that invokes UpdateTeamAdmin on the service.
func MakeUpdateTeamAdminEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateTeamAdminRequest)
		e0, m1 := s.UpdateTeamAdmin(ctx, req.TeamId, req.AdminId, req.Admin)
		return UpdateTeamAdminResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateTeamAdminResponse) Failed() error {
	return r.E0
}

// UpdateTeamMemberRequest collects the request parameters for the UpdateTeamMember method.
type UpdateTeamMemberRequest struct {
	TeamId   int32      `json:"team_id"`
	MemberId int32      `json:"member_id"`
	Member   proto.User `json:"member"`
}

// UpdateTeamMemberResponse collects the response parameters for the UpdateTeamMember method.
type UpdateTeamMemberResponse struct {
	E0 error      `json:"e0"`
	M1 proto.Team `json:"m1"`
}

// MakeUpdateTeamMemberEndpoint returns an endpoint that invokes UpdateTeamMember on the service.
func MakeUpdateTeamMemberEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateTeamMemberRequest)
		e0, m1 := s.UpdateTeamMember(ctx, req.TeamId, req.MemberId, req.Member)
		return UpdateTeamMemberResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateTeamMemberResponse) Failed() error {
	return r.E0
}

// UpdateTeamAdvisorRequest collects the request parameters for the UpdateTeamAdvisor method.
type UpdateTeamAdvisorRequest struct {
	TeamId    int32      `json:"team_id"`
	AdvisorId int32      `json:"advisor_id"`
	Advisor   proto.User `json:"advisor"`
}

// UpdateTeamAdvisorResponse collects the response parameters for the UpdateTeamAdvisor method.
type UpdateTeamAdvisorResponse struct {
	E0 error      `json:"e0"`
	M1 proto.Team `json:"m1"`
}

// MakeUpdateTeamAdvisorEndpoint returns an endpoint that invokes UpdateTeamAdvisor on the service.
func MakeUpdateTeamAdvisorEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateTeamAdvisorRequest)
		e0, m1 := s.UpdateTeamAdvisor(ctx, req.TeamId, req.AdvisorId, req.Advisor)
		return UpdateTeamAdvisorResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateTeamAdvisorResponse) Failed() error {
	return r.E0
}

// UpdateTeamProfileRequest collects the request parameters for the UpdateTeamProfile method.
type UpdateTeamProfileRequest struct {
	TeamId    int32             `json:"team_id"`
	ProfileId int32             `json:"profile_id"`
	Profile   proto.TeamProfile `json:"profile"`
}

// UpdateTeamProfileResponse collects the response parameters for the UpdateTeamProfile method.
type UpdateTeamProfileResponse struct {
	E0 error      `json:"e0"`
	M1 proto.Team `json:"m1"`
}

// MakeUpdateTeamProfileEndpoint returns an endpoint that invokes UpdateTeamProfile on the service.
func MakeUpdateTeamProfileEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateTeamProfileRequest)
		e0, m1 := s.UpdateTeamProfile(ctx, req.TeamId, req.ProfileId, req.Profile)
		return UpdateTeamProfileResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateTeamProfileResponse) Failed() error {
	return r.E0
}

// UpdateTeamSubscriptionRequest collects the request parameters for the UpdateTeamSubscription method.
type UpdateTeamSubscriptionRequest struct {
	TeamId         int32               `json:"team_id"`
	SubscriptionId int32               `json:"subscription_id"`
	Subcription    proto.Subscriptions `json:"subcription"`
}

// UpdateTeamSubscriptionResponse collects the response parameters for the UpdateTeamSubscription method.
type UpdateTeamSubscriptionResponse struct {
	E0 error      `json:"e0"`
	M1 proto.Team `json:"m1"`
}

// MakeUpdateTeamSubscriptionEndpoint returns an endpoint that invokes UpdateTeamSubscription on the service.
func MakeUpdateTeamSubscriptionEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateTeamSubscriptionRequest)
		e0, m1 := s.UpdateTeamSubscription(ctx, req.TeamId, req.SubscriptionId, req.Subcription)
		return UpdateTeamSubscriptionResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateTeamSubscriptionResponse) Failed() error {
	return r.E0
}

// DeleteTeamRequest collects the request parameters for the DeleteTeam method.
type DeleteTeamRequest struct {
	TeamId int32 `json:"team_id"`
}

// DeleteTeamResponse collects the response parameters for the DeleteTeam method.
type DeleteTeamResponse struct {
	E0 error `json:"e0"`
}

// MakeDeleteTeamEndpoint returns an endpoint that invokes DeleteTeam on the service.
func MakeDeleteTeamEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteTeamRequest)
		e0 := s.DeleteTeam(ctx, req.TeamId)
		return DeleteTeamResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r DeleteTeamResponse) Failed() error {
	return r.E0
}

// DeleteTeamProfileRequest collects the request parameters for the DeleteTeamProfile method.
type DeleteTeamProfileRequest struct {
	TeamId    int32 `json:"team_id"`
	ProfileId int32 `json:"profile_id"`
}

// DeleteTeamProfileResponse collects the response parameters for the DeleteTeamProfile method.
type DeleteTeamProfileResponse struct {
	E0 error `json:"e0"`
}

// MakeDeleteTeamProfileEndpoint returns an endpoint that invokes DeleteTeamProfile on the service.
func MakeDeleteTeamProfileEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteTeamProfileRequest)
		e0 := s.DeleteTeamProfile(ctx, req.TeamId, req.ProfileId)
		return DeleteTeamProfileResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r DeleteTeamProfileResponse) Failed() error {
	return r.E0
}

// DeleteTeamMemberRequest collects the request parameters for the DeleteTeamMember method.
type DeleteTeamMemberRequest struct {
	TeamId       int32 `json:"team_id"`
	TeamMemberId int32 `json:"team_member_id"`
}

// DeleteTeamMemberResponse collects the response parameters for the DeleteTeamMember method.
type DeleteTeamMemberResponse struct {
	E0 error `json:"e0"`
}

// MakeDeleteTeamMemberEndpoint returns an endpoint that invokes DeleteTeamMember on the service.
func MakeDeleteTeamMemberEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteTeamMemberRequest)
		e0 := s.DeleteTeamMember(ctx, req.TeamId, req.TeamMemberId)
		return DeleteTeamMemberResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r DeleteTeamMemberResponse) Failed() error {
	return r.E0
}

// DeleteTeamAdvisorRequest collects the request parameters for the DeleteTeamAdvisor method.
type DeleteTeamAdvisorRequest struct {
	TeamId    int32 `json:"team_id"`
	AdvisorId int32 `json:"advisor_id"`
}

// DeleteTeamAdvisorResponse collects the response parameters for the DeleteTeamAdvisor method.
type DeleteTeamAdvisorResponse struct {
	E0 error `json:"e0"`
}

// MakeDeleteTeamAdvisorEndpoint returns an endpoint that invokes DeleteTeamAdvisor on the service.
func MakeDeleteTeamAdvisorEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteTeamAdvisorRequest)
		e0 := s.DeleteTeamAdvisor(ctx, req.TeamId, req.AdvisorId)
		return DeleteTeamAdvisorResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r DeleteTeamAdvisorResponse) Failed() error {
	return r.E0
}

// DeleteTeamAdminRequest collects the request parameters for the DeleteTeamAdmin method.
type DeleteTeamAdminRequest struct {
	TeamId  int32 `json:"team_id"`
	AdminId int32 `json:"admin_id"`
}

// DeleteTeamAdminResponse collects the response parameters for the DeleteTeamAdmin method.
type DeleteTeamAdminResponse struct {
	E0 error `json:"e0"`
}

// MakeDeleteTeamAdminEndpoint returns an endpoint that invokes DeleteTeamAdmin on the service.
func MakeDeleteTeamAdminEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteTeamAdminRequest)
		e0 := s.DeleteTeamAdmin(ctx, req.TeamId, req.AdminId)
		return DeleteTeamAdminResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r DeleteTeamAdminResponse) Failed() error {
	return r.E0
}

// DeleteTeamSubscriptionRequest collects the request parameters for the DeleteTeamSubscription method.
type DeleteTeamSubscriptionRequest struct {
	TeamId         int32 `json:"team_id"`
	SubscriptionId int32 `json:"subscription_id"`
}

// DeleteTeamSubscriptionResponse collects the response parameters for the DeleteTeamSubscription method.
type DeleteTeamSubscriptionResponse struct {
	E0 error `json:"e0"`
}

// MakeDeleteTeamSubscriptionEndpoint returns an endpoint that invokes DeleteTeamSubscription on the service.
func MakeDeleteTeamSubscriptionEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteTeamSubscriptionRequest)
		e0 := s.DeleteTeamSubscription(ctx, req.TeamId, req.SubscriptionId)
		return DeleteTeamSubscriptionResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r DeleteTeamSubscriptionResponse) Failed() error {
	return r.E0
}

// GetTeamRequest collects the request parameters for the GetTeam method.
type GetTeamRequest struct {
	TeamId int32 `json:"team_id"`
}

// GetTeamResponse collects the response parameters for the GetTeam method.
type GetTeamResponse struct {
	E0 error      `json:"e0"`
	M1 proto.Team `json:"m1"`
}

// MakeGetTeamEndpoint returns an endpoint that invokes GetTeam on the service.
func MakeGetTeamEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTeamRequest)
		e0, m1 := s.GetTeam(ctx, req.TeamId)
		return GetTeamResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetTeamResponse) Failed() error {
	return r.E0
}

// GetTeamProfileRequest collects the request parameters for the GetTeamProfile method.
type GetTeamProfileRequest struct {
	TeamId int32 `json:"team_id"`
}

// GetTeamProfileResponse collects the response parameters for the GetTeamProfile method.
type GetTeamProfileResponse struct {
	E0 error         `json:"e0"`
	M1 proto.Profile `json:"m1"`
}

// MakeGetTeamProfileEndpoint returns an endpoint that invokes GetTeamProfile on the service.
func MakeGetTeamProfileEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTeamProfileRequest)
		e0, m1 := s.GetTeamProfile(ctx, req.TeamId)
		return GetTeamProfileResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetTeamProfileResponse) Failed() error {
	return r.E0
}

// GetTeamSubscriptionsRequest collects the request parameters for the GetTeamSubscriptions method.
type GetTeamSubscriptionsRequest struct {
	TeamId int32 `json:"team_id"`
}

// GetTeamSubscriptionsResponse collects the response parameters for the GetTeamSubscriptions method.
type GetTeamSubscriptionsResponse struct {
	E0 error                 `json:"e0"`
	M1 []model.Subscriptions `json:"m1"`
}

// MakeGetTeamSubscriptionsEndpoint returns an endpoint that invokes GetTeamSubscriptions on the service.
func MakeGetTeamSubscriptionsEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTeamSubscriptionsRequest)
		e0, m1 := s.GetTeamSubscriptions(ctx, req.TeamId)
		return GetTeamSubscriptionsResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetTeamSubscriptionsResponse) Failed() error {
	return r.E0
}

// GetTeamMembersRequest collects the request parameters for the GetTeamMembers method.
type GetTeamMembersRequest struct {
	TeamId int32 `json:"team_id"`
}

// GetTeamMembersResponse collects the response parameters for the GetTeamMembers method.
type GetTeamMembersResponse struct {
	E0 error        `json:"e0"`
	M1 []model.User `json:"m1"`
}

// MakeGetTeamMembersEndpoint returns an endpoint that invokes GetTeamMembers on the service.
func MakeGetTeamMembersEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTeamMembersRequest)
		e0, m1 := s.GetTeamMembers(ctx, req.TeamId)
		return GetTeamMembersResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetTeamMembersResponse) Failed() error {
	return r.E0
}

// GetTeamAdvisorsRequest collects the request parameters for the GetTeamAdvisors method.
type GetTeamAdvisorsRequest struct {
	TeamId int32 `json:"team_id"`
}

// GetTeamAdvisorsResponse collects the response parameters for the GetTeamAdvisors method.
type GetTeamAdvisorsResponse struct {
	E0 error        `json:"e0"`
	M1 []model.User `json:"m1"`
}

// MakeGetTeamAdvisorsEndpoint returns an endpoint that invokes GetTeamAdvisors on the service.
func MakeGetTeamAdvisorsEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTeamAdvisorsRequest)
		e0, m1 := s.GetTeamAdvisors(ctx, req.TeamId)
		return GetTeamAdvisorsResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetTeamAdvisorsResponse) Failed() error {
	return r.E0
}

// GetTeamAdminRequest collects the request parameters for the GetTeamAdmin method.
type GetTeamAdminRequest struct {
	TeamId int32 `json:"team_id"`
}

// GetTeamAdminResponse collects the response parameters for the GetTeamAdmin method.
type GetTeamAdminResponse struct {
	E0 error      `json:"e0"`
	M1 proto.User `json:"m1"`
}

// MakeGetTeamAdminEndpoint returns an endpoint that invokes GetTeamAdmin on the service.
func MakeGetTeamAdminEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTeamAdminRequest)
		e0, m1 := s.GetTeamAdmin(ctx, req.TeamId)
		return GetTeamAdminResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetTeamAdminResponse) Failed() error {
	return r.E0
}

// GetTeamsRequest collects the request parameters for the GetTeams method.
type GetTeamsRequest struct {
	Limit int32 `json:"limit"`
}

// GetTeamsResponse collects the response parameters for the GetTeams method.
type GetTeamsResponse struct {
	E0 error        `json:"e0"`
	M1 []model.Team `json:"m1"`
}

// MakeGetTeamsEndpoint returns an endpoint that invokes GetTeams on the service.
func MakeGetTeamsEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTeamsRequest)
		e0, m1 := s.GetTeams(ctx, req.Limit)
		return GetTeamsResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetTeamsResponse) Failed() error {
	return r.E0
}

// GetTeamsByTypeRequest collects the request parameters for the GetTeamsByType method.
type GetTeamsByTypeRequest struct {
	TeamType string `json:"team_type"`
	Limit    int32  `json:"limit"`
}

// GetTeamsByTypeResponse collects the response parameters for the GetTeamsByType method.
type GetTeamsByTypeResponse struct {
	E0 error        `json:"e0"`
	M1 []model.Team `json:"m1"`
}

// MakeGetTeamsByTypeEndpoint returns an endpoint that invokes GetTeamsByType on the service.
func MakeGetTeamsByTypeEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTeamsByTypeRequest)
		e0, m1 := s.GetTeamsByType(ctx, req.TeamType, req.Limit)
		return GetTeamsByTypeResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetTeamsByTypeResponse) Failed() error {
	return r.E0
}

// GetTeamsByIndustryRequest collects the request parameters for the GetTeamsByIndustry method.
type GetTeamsByIndustryRequest struct {
	Industry string `json:"industry"`
	Limit    int32  `json:"limit"`
}

// GetTeamsByIndustryResponse collects the response parameters for the GetTeamsByIndustry method.
type GetTeamsByIndustryResponse struct {
	E0 error        `json:"e0"`
	M1 []model.Team `json:"m1"`
}

// MakeGetTeamsByIndustryEndpoint returns an endpoint that invokes GetTeamsByIndustry on the service.
func MakeGetTeamsByIndustryEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTeamsByIndustryRequest)
		e0, m1 := s.GetTeamsByIndustry(ctx, req.Industry, req.Limit)
		return GetTeamsByIndustryResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetTeamsByIndustryResponse) Failed() error {
	return r.E0
}

// GetTeamsByNumberOfEmployeesRequest collects the request parameters for the GetTeamsByNumberOfEmployees method.
type GetTeamsByNumberOfEmployeesRequest struct {
	NumEmployees int32 `json:"num_employees"`
	Limit        int32 `json:"limit"`
}

// GetTeamsByNumberOfEmployeesResponse collects the response parameters for the GetTeamsByNumberOfEmployees method.
type GetTeamsByNumberOfEmployeesResponse struct {
	E0 error        `json:"e0"`
	M1 []model.Team `json:"m1"`
}

// MakeGetTeamsByNumberOfEmployeesEndpoint returns an endpoint that invokes GetTeamsByNumberOfEmployees on the service.
func MakeGetTeamsByNumberOfEmployeesEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTeamsByNumberOfEmployeesRequest)
		e0, m1 := s.GetTeamsByNumberOfEmployees(ctx, req.NumEmployees, req.Limit)
		return GetTeamsByNumberOfEmployeesResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetTeamsByNumberOfEmployeesResponse) Failed() error {
	return r.E0
}

// GetTeamProfilesRequest collects the request parameters for the GetTeamProfiles method.
type GetTeamProfilesRequest struct {
	Limit int32 `json:"limit"`
}

// GetTeamProfilesResponse collects the response parameters for the GetTeamProfiles method.
type GetTeamProfilesResponse struct {
	E0 error               `json:"e0"`
	M1 []model.TeamProfile `json:"m1"`
}

// MakeGetTeamProfilesEndpoint returns an endpoint that invokes GetTeamProfiles on the service.
func MakeGetTeamProfilesEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTeamProfilesRequest)
		e0, m1 := s.GetTeamProfiles(ctx, req.Limit)
		return GetTeamProfilesResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetTeamProfilesResponse) Failed() error {
	return r.E0
}

// GetTeamsByTagsRequest collects the request parameters for the GetTeamsByTags method.
type GetTeamsByTagsRequest struct {
	Tags  string `json:"tags"`
	Limit int32  `json:"limit"`
}

// GetTeamsByTagsResponse collects the response parameters for the GetTeamsByTags method.
type GetTeamsByTagsResponse struct {
	E0 error               `json:"e0"`
	M1 []model.TeamProfile `json:"m1"`
}

// MakeGetTeamsByTagsEndpoint returns an endpoint that invokes GetTeamsByTags on the service.
func MakeGetTeamsByTagsEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTeamsByTagsRequest)
		e0, m1 := s.GetTeamsByTags(ctx, req.Tags, req.Limit)
		return GetTeamsByTagsResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetTeamsByTagsResponse) Failed() error {
	return r.E0
}

// CreateGroupRequest collects the request parameters for the CreateGroup method.
type CreateGroupRequest struct {
	Group proto.Group `json:"group"`
	Admin proto.User  `json:"admin"`
}

// CreateGroupResponse collects the response parameters for the CreateGroup method.
type CreateGroupResponse struct {
	E0 error `json:"e0"`
}

// MakeCreateGroupEndpoint returns an endpoint that invokes CreateGroup on the service.
func MakeCreateGroupEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateGroupRequest)
		e0 := s.CreateGroup(ctx, req.Group, req.Admin)
		return CreateGroupResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r CreateGroupResponse) Failed() error {
	return r.E0
}

// AddGroupMemberRequest collects the request parameters for the AddGroupMember method.
type AddGroupMemberRequest struct {
	GroupId int32      `json:"group_id"`
	Member  proto.User `json:"member"`
}

// AddGroupMemberResponse collects the response parameters for the AddGroupMember method.
type AddGroupMemberResponse struct {
	E0 error `json:"e0"`
}

// MakeAddGroupMemberEndpoint returns an endpoint that invokes AddGroupMember on the service.
func MakeAddGroupMemberEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddGroupMemberRequest)
		e0 := s.AddGroupMember(ctx, req.GroupId, req.Member)
		return AddGroupMemberResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r AddGroupMemberResponse) Failed() error {
	return r.E0
}

// UpdateGroupRequest collects the request parameters for the UpdateGroup method.
type UpdateGroupRequest struct {
	GroupId int32       `json:"group_id"`
	Group   proto.Group `json:"group"`
}

// UpdateGroupResponse collects the response parameters for the UpdateGroup method.
type UpdateGroupResponse struct {
	E0 error `json:"e0"`
}

// MakeUpdateGroupEndpoint returns an endpoint that invokes UpdateGroup on the service.
func MakeUpdateGroupEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateGroupRequest)
		e0 := s.UpdateGroup(ctx, req.GroupId, req.Group)
		return UpdateGroupResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r UpdateGroupResponse) Failed() error {
	return r.E0
}

// UpdateGroupMemberRequest collects the request parameters for the UpdateGroupMember method.
type UpdateGroupMemberRequest struct {
	GroupId       int32      `json:"group_id"`
	GroupMemberId int32      `json:"group_member_id"`
	Member        proto.User `json:"member"`
}

// UpdateGroupMemberResponse collects the response parameters for the UpdateGroupMember method.
type UpdateGroupMemberResponse struct {
	E0 error `json:"e0"`
}

// MakeUpdateGroupMemberEndpoint returns an endpoint that invokes UpdateGroupMember on the service.
func MakeUpdateGroupMemberEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateGroupMemberRequest)
		e0 := s.UpdateGroupMember(ctx, req.GroupId, req.GroupMemberId, req.Member)
		return UpdateGroupMemberResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r UpdateGroupMemberResponse) Failed() error {
	return r.E0
}

// DeleteGroupRequest collects the request parameters for the DeleteGroup method.
type DeleteGroupRequest struct {
	GroupId int32 `json:"group_id"`
}

// DeleteGroupResponse collects the response parameters for the DeleteGroup method.
type DeleteGroupResponse struct {
	E0 error `json:"e0"`
}

// MakeDeleteGroupEndpoint returns an endpoint that invokes DeleteGroup on the service.
func MakeDeleteGroupEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteGroupRequest)
		e0 := s.DeleteGroup(ctx, req.GroupId)
		return DeleteGroupResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r DeleteGroupResponse) Failed() error {
	return r.E0
}

// DeleteGroupMemberRequest collects the request parameters for the DeleteGroupMember method.
type DeleteGroupMemberRequest struct {
	GroupId  int32 `json:"group_id"`
	MemberId int32 `json:"member_id"`
}

// DeleteGroupMemberResponse collects the response parameters for the DeleteGroupMember method.
type DeleteGroupMemberResponse struct {
	E0 error `json:"e0"`
}

// MakeDeleteGroupMemberEndpoint returns an endpoint that invokes DeleteGroupMember on the service.
func MakeDeleteGroupMemberEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteGroupMemberRequest)
		e0 := s.DeleteGroupMember(ctx, req.GroupId, req.MemberId)
		return DeleteGroupMemberResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r DeleteGroupMemberResponse) Failed() error {
	return r.E0
}

// GetGroupRequest collects the request parameters for the GetGroup method.
type GetGroupRequest struct {
	GroupId int32 `json:"group_id"`
}

// GetGroupResponse collects the response parameters for the GetGroup method.
type GetGroupResponse struct {
	E0 error       `json:"e0"`
	M1 proto.Group `json:"m1"`
}

// MakeGetGroupEndpoint returns an endpoint that invokes GetGroup on the service.
func MakeGetGroupEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetGroupRequest)
		e0, m1 := s.GetGroup(ctx, req.GroupId)
		return GetGroupResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetGroupResponse) Failed() error {
	return r.E0
}

// GetGroupsRequest collects the request parameters for the GetGroups method.
type GetGroupsRequest struct {
	Limit int32 `json:"limit"`
}

// GetGroupsResponse collects the response parameters for the GetGroups method.
type GetGroupsResponse struct {
	E0 error         `json:"e0"`
	M1 []model.Group `json:"m1"`
}

// MakeGetGroupsEndpoint returns an endpoint that invokes GetGroups on the service.
func MakeGetGroupsEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetGroupsRequest)
		e0, m1 := s.GetGroups(ctx, req.Limit)
		return GetGroupsResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetGroupsResponse) Failed() error {
	return r.E0
}

// GetGroupsByTypeRequest collects the request parameters for the GetGroupsByType method.
type GetGroupsByTypeRequest struct {
	GroupType string `json:"group_type"`
	Limit     int32  `json:"limit"`
}

// GetGroupsByTypeResponse collects the response parameters for the GetGroupsByType method.
type GetGroupsByTypeResponse struct {
	E0 error         `json:"e0"`
	M1 []model.Group `json:"m1"`
}

// MakeGetGroupsByTypeEndpoint returns an endpoint that invokes GetGroupsByType on the service.
func MakeGetGroupsByTypeEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetGroupsByTypeRequest)
		e0, m1 := s.GetGroupsByType(ctx, req.GroupType, req.Limit)
		return GetGroupsByTypeResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetGroupsByTypeResponse) Failed() error {
	return r.E0
}

// GetGroupsByNumMembersRequest collects the request parameters for the GetGroupsByNumMembers method.
type GetGroupsByNumMembersRequest struct {
	MaxNumEmployees int32 `json:"max_num_employees"`
	Limit           int32 `json:"limit"`
}

// GetGroupsByNumMembersResponse collects the response parameters for the GetGroupsByNumMembers method.
type GetGroupsByNumMembersResponse struct {
	E0 error         `json:"e0"`
	M1 []model.Group `json:"m1"`
}

// MakeGetGroupsByNumMembersEndpoint returns an endpoint that invokes GetGroupsByNumMembers on the service.
func MakeGetGroupsByNumMembersEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetGroupsByNumMembersRequest)
		e0, m1 := s.GetGroupsByNumMembers(ctx, req.MaxNumEmployees, req.Limit)
		return GetGroupsByNumMembersResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetGroupsByNumMembersResponse) Failed() error {
	return r.E0
}

// GetGroupsByTagsRequest collects the request parameters for the GetGroupsByTags method.
type GetGroupsByTagsRequest struct {
	Tags  string `json:"tags"`
	Limit int32  `json:"limit"`
}

// GetGroupsByTagsResponse collects the response parameters for the GetGroupsByTags method.
type GetGroupsByTagsResponse struct {
	E0 error         `json:"e0"`
	M1 []model.Group `json:"m1"`
}

// MakeGetGroupsByTagsEndpoint returns an endpoint that invokes GetGroupsByTags on the service.
func MakeGetGroupsByTagsEndpoint(s service.MicroService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetGroupsByTagsRequest)
		e0, m1 := s.GetGroupsByTags(ctx, req.Tags, req.Limit)
		return GetGroupsByTagsResponse{
			E0: e0,
			M1: m1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetGroupsByTagsResponse) Failed() error {
	return r.E0
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateUser implements Service. Primarily useful in a client.
func (e Endpoints) CreateUser(ctx context.Context, user proto.User) (e0 error) {
	request := CreateUserRequest{User: user}
	response, err := e.CreateUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateUserResponse).E0
}

// CreateProfile implements Service. Primarily useful in a client.
func (e Endpoints) CreateProfile(ctx context.Context, user_id int32, profile proto.Profile) (e0 error) {
	request := CreateProfileRequest{
		Profile: profile,
		UserId:  user_id,
	}
	response, err := e.CreateProfileEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateProfileResponse).E0
}

// CreateSubscription implements Service. Primarily useful in a client.
func (e Endpoints) CreateSubscription(ctx context.Context, user_id int32, subscription proto.Subscriptions) (e0 error) {
	request := CreateSubscriptionRequest{
		Subscription: subscription,
		UserId:       user_id,
	}
	response, err := e.CreateSubscriptionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateSubscriptionResponse).E0
}

// CreateSubscriptions implements Service. Primarily useful in a client.
func (e Endpoints) CreateSubscriptions(ctx context.Context, user_id []model.Subscriptions, subscriptions []model.Subscriptions) (e0 error) {
	request := CreateSubscriptionsRequest{
		Subscriptions: subscriptions,
		UserId:        user_id,
	}
	response, err := e.CreateSubscriptionsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateSubscriptionsResponse).E0
}

// UpdateUser implements Service. Primarily useful in a client.
func (e Endpoints) UpdateUser(ctx context.Context, user_id int32, user proto.User) (e0 error, m1 proto.User) {
	request := UpdateUserRequest{
		User:   user,
		UserId: user_id,
	}
	response, err := e.UpdateUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateUserResponse).E0, response.(UpdateUserResponse).M1
}

// UpdateUserProfile implements Service. Primarily useful in a client.
func (e Endpoints) UpdateUserProfile(ctx context.Context, user_id int32, profile_id int32, profile proto.Profile) (e0 error, m1 proto.Profile) {
	request := UpdateUserProfileRequest{
		Profile:   profile,
		ProfileId: profile_id,
		UserId:    user_id,
	}
	response, err := e.UpdateUserProfileEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateUserProfileResponse).E0, response.(UpdateUserProfileResponse).M1
}

// UpdateUserSubscription implements Service. Primarily useful in a client.
func (e Endpoints) UpdateUserSubscription(ctx context.Context, user_id int32, subscription_id int32, subscription proto.Subscriptions) (e0 error, m1 proto.Subscriptions) {
	request := UpdateUserSubscriptionRequest{
		Subscription:   subscription,
		SubscriptionId: subscription_id,
		UserId:         user_id,
	}
	response, err := e.UpdateUserSubscriptionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateUserSubscriptionResponse).E0, response.(UpdateUserSubscriptionResponse).M1
}

// DeleteUser implements Service. Primarily useful in a client.
func (e Endpoints) DeleteUser(ctx context.Context, user_id int32) (e0 error) {
	request := DeleteUserRequest{UserId: user_id}
	response, err := e.DeleteUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteUserResponse).E0
}

// DeleteUserProfile implements Service. Primarily useful in a client.
func (e Endpoints) DeleteUserProfile(ctx context.Context, user_id int32, profile_id int32) (e0 error) {
	request := DeleteUserProfileRequest{
		ProfileId: profile_id,
		UserId:    user_id,
	}
	response, err := e.DeleteUserProfileEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteUserProfileResponse).E0
}

// DeleteSubscription implements Service. Primarily useful in a client.
func (e Endpoints) DeleteSubscription(ctx context.Context, user_id int32, subscription_id int32) (e0 error) {
	request := DeleteSubscriptionRequest{
		SubscriptionId: subscription_id,
		UserId:         user_id,
	}
	response, err := e.DeleteSubscriptionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteSubscriptionResponse).E0
}

// GetUser implements Service. Primarily useful in a client.
func (e Endpoints) GetUser(ctx context.Context, user_id int32) (e0 error, m1 proto.User) {
	request := GetUserRequest{UserId: user_id}
	response, err := e.GetUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserResponse).E0, response.(GetUserResponse).M1
}

// GetUsers implements Service. Primarily useful in a client.
func (e Endpoints) GetUsers(ctx context.Context, limit int32) (e0 error, m1 []model.User) {
	request := GetUsersRequest{Limit: limit}
	response, err := e.GetUsersEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUsersResponse).E0, response.(GetUsersResponse).M1
}

// GetUsersByAccountType implements Service. Primarily useful in a client.
func (e Endpoints) GetUsersByAccountType(ctx context.Context, accounType string, limit int32) (e0 error, m1 []model.User) {
	request := GetUsersByAccountTypeRequest{
		AccounType: accounType,
		Limit:      limit,
	}
	response, err := e.GetUsersByAccountTypeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUsersByAccountTypeResponse).E0, response.(GetUsersByAccountTypeResponse).M1
}

// GetUsersByIntent implements Service. Primarily useful in a client.
func (e Endpoints) GetUsersByIntent(ctx context.Context, intent string, limit int32) (e0 error, m1 []model.User) {
	request := GetUsersByIntentRequest{
		Intent: intent,
		Limit:  limit,
	}
	response, err := e.GetUsersByIntentEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUsersByIntentResponse).E0, response.(GetUsersByIntentResponse).M1
}

// GetUserProfile implements Service. Primarily useful in a client.
func (e Endpoints) GetUserProfile(ctx context.Context, user_id int32) (e0 error, m1 proto.Profile) {
	request := GetUserProfileRequest{UserId: user_id}
	response, err := e.GetUserProfileEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserProfileResponse).E0, response.(GetUserProfileResponse).M1
}

// GetUserProfiles implements Service. Primarily useful in a client.
func (e Endpoints) GetUserProfiles(ctx context.Context, limit int32) (e0 error, m1 []model.Profile) {
	request := GetUserProfilesRequest{Limit: limit}
	response, err := e.GetUserProfilesEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserProfilesResponse).E0, response.(GetUserProfilesResponse).M1
}

// GetUserProfilesByType implements Service. Primarily useful in a client.
func (e Endpoints) GetUserProfilesByType(ctx context.Context, profileType string, limit int32) (e0 error, m1 []model.Profile) {
	request := GetUserProfilesByTypeRequest{
		Limit:       limit,
		ProfileType: profileType,
	}
	response, err := e.GetUserProfilesByTypeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserProfilesByTypeResponse).E0, response.(GetUserProfilesByTypeResponse).M1
}

// GetUserProfilesByNationality implements Service. Primarily useful in a client.
func (e Endpoints) GetUserProfilesByNationality(ctx context.Context, nationality string, limit int32) (e0 error, m1 []model.Profile) {
	request := GetUserProfilesByNationalityRequest{
		Limit:       limit,
		Nationality: nationality,
	}
	response, err := e.GetUserProfilesByNationalityEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserProfilesByNationalityResponse).E0, response.(GetUserProfilesByNationalityResponse).M1
}

// GetUserSubscriptions implements Service. Primarily useful in a client.
func (e Endpoints) GetUserSubscriptions(ctx context.Context, user_id int32) (e0 error, m1 []model.Subscriptions) {
	request := GetUserSubscriptionsRequest{UserId: user_id}
	response, err := e.GetUserSubscriptionsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserSubscriptionsResponse).E0, response.(GetUserSubscriptionsResponse).M1
}

// CreateTeam implements Service. Primarily useful in a client.
func (e Endpoints) CreateTeam(ctx context.Context, team proto.Team, admin proto.User) (e0 error) {
	request := CreateTeamRequest{
		Admin: admin,
		Team:  team,
	}
	response, err := e.CreateTeamEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateTeamResponse).E0
}

// CreatTeamProfile implements Service. Primarily useful in a client.
func (e Endpoints) CreatTeamProfile(ctx context.Context, team_id int32, profile proto.TeamProfile) (e0 error) {
	request := CreatTeamProfileRequest{
		Profile: profile,
		TeamId:  team_id,
	}
	response, err := e.CreatTeamProfileEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreatTeamProfileResponse).E0
}

// CreateTeamSubscription implements Service. Primarily useful in a client.
func (e Endpoints) CreateTeamSubscription(ctx context.Context, team_id int32, subscription proto.Subscriptions) (e0 error) {
	request := CreateTeamSubscriptionRequest{
		Subscription: subscription,
		TeamId:       team_id,
	}
	response, err := e.CreateTeamSubscriptionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateTeamSubscriptionResponse).E0
}

// AddMemberToTeam implements Service. Primarily useful in a client.
func (e Endpoints) AddMemberToTeam(ctx context.Context, team_id int32, member proto.User) (e0 error) {
	request := AddMemberToTeamRequest{
		Member: member,
		TeamId: team_id,
	}
	response, err := e.AddMemberToTeamEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddMemberToTeamResponse).E0
}

// AddAdvisorToTeam implements Service. Primarily useful in a client.
func (e Endpoints) AddAdvisorToTeam(ctx context.Context, team_id int32, advisor_id int32, advisor proto.User) (e0 error) {
	request := AddAdvisorToTeamRequest{
		Advisor:   advisor,
		AdvisorId: advisor_id,
		TeamId:    team_id,
	}
	response, err := e.AddAdvisorToTeamEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddAdvisorToTeamResponse).E0
}

// UpdateTeam implements Service. Primarily useful in a client.
func (e Endpoints) UpdateTeam(ctx context.Context, team_id int32, team proto.Team) (e0 error, m1 proto.Team) {
	request := UpdateTeamRequest{
		Team:   team,
		TeamId: team_id,
	}
	response, err := e.UpdateTeamEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateTeamResponse).E0, response.(UpdateTeamResponse).M1
}

// UpdateTeamAdmin implements Service. Primarily useful in a client.
func (e Endpoints) UpdateTeamAdmin(ctx context.Context, team_id int32, admin_id int32, admin proto.User) (e0 error, m1 proto.Team) {
	request := UpdateTeamAdminRequest{
		Admin:   admin,
		AdminId: admin_id,
		TeamId:  team_id,
	}
	response, err := e.UpdateTeamAdminEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateTeamAdminResponse).E0, response.(UpdateTeamAdminResponse).M1
}

// UpdateTeamMember implements Service. Primarily useful in a client.
func (e Endpoints) UpdateTeamMember(ctx context.Context, team_id int32, member_id int32, member proto.User) (e0 error, m1 proto.Team) {
	request := UpdateTeamMemberRequest{
		Member:   member,
		MemberId: member_id,
		TeamId:   team_id,
	}
	response, err := e.UpdateTeamMemberEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateTeamMemberResponse).E0, response.(UpdateTeamMemberResponse).M1
}

// UpdateTeamAdvisor implements Service. Primarily useful in a client.
func (e Endpoints) UpdateTeamAdvisor(ctx context.Context, team_id int32, advisor_id int32, advisor proto.User) (e0 error, m1 proto.Team) {
	request := UpdateTeamAdvisorRequest{
		Advisor:   advisor,
		AdvisorId: advisor_id,
		TeamId:    team_id,
	}
	response, err := e.UpdateTeamAdvisorEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateTeamAdvisorResponse).E0, response.(UpdateTeamAdvisorResponse).M1
}

// UpdateTeamProfile implements Service. Primarily useful in a client.
func (e Endpoints) UpdateTeamProfile(ctx context.Context, team_id int32, profile_id int32, profile proto.TeamProfile) (e0 error, m1 proto.Team) {
	request := UpdateTeamProfileRequest{
		Profile:   profile,
		ProfileId: profile_id,
		TeamId:    team_id,
	}
	response, err := e.UpdateTeamProfileEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateTeamProfileResponse).E0, response.(UpdateTeamProfileResponse).M1
}

// UpdateTeamSubscription implements Service. Primarily useful in a client.
func (e Endpoints) UpdateTeamSubscription(ctx context.Context, team_id int32, subscription_id int32, subcription proto.Subscriptions) (e0 error, m1 proto.Team) {
	request := UpdateTeamSubscriptionRequest{
		Subcription:    subcription,
		SubscriptionId: subscription_id,
		TeamId:         team_id,
	}
	response, err := e.UpdateTeamSubscriptionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateTeamSubscriptionResponse).E0, response.(UpdateTeamSubscriptionResponse).M1
}

// DeleteTeam implements Service. Primarily useful in a client.
func (e Endpoints) DeleteTeam(ctx context.Context, team_id int32) (e0 error) {
	request := DeleteTeamRequest{TeamId: team_id}
	response, err := e.DeleteTeamEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteTeamResponse).E0
}

// DeleteTeamProfile implements Service. Primarily useful in a client.
func (e Endpoints) DeleteTeamProfile(ctx context.Context, team_id int32, profile_id int32) (e0 error) {
	request := DeleteTeamProfileRequest{
		ProfileId: profile_id,
		TeamId:    team_id,
	}
	response, err := e.DeleteTeamProfileEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteTeamProfileResponse).E0
}

// DeleteTeamMember implements Service. Primarily useful in a client.
func (e Endpoints) DeleteTeamMember(ctx context.Context, team_id int32, team_member_id int32) (e0 error) {
	request := DeleteTeamMemberRequest{
		TeamId:       team_id,
		TeamMemberId: team_member_id,
	}
	response, err := e.DeleteTeamMemberEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteTeamMemberResponse).E0
}

// DeleteTeamAdvisor implements Service. Primarily useful in a client.
func (e Endpoints) DeleteTeamAdvisor(ctx context.Context, team_id int32, advisor_id int32) (e0 error) {
	request := DeleteTeamAdvisorRequest{
		AdvisorId: advisor_id,
		TeamId:    team_id,
	}
	response, err := e.DeleteTeamAdvisorEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteTeamAdvisorResponse).E0
}

// DeleteTeamAdmin implements Service. Primarily useful in a client.
func (e Endpoints) DeleteTeamAdmin(ctx context.Context, team_id int32, admin_id int32) (e0 error) {
	request := DeleteTeamAdminRequest{
		AdminId: admin_id,
		TeamId:  team_id,
	}
	response, err := e.DeleteTeamAdminEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteTeamAdminResponse).E0
}

// DeleteTeamSubscription implements Service. Primarily useful in a client.
func (e Endpoints) DeleteTeamSubscription(ctx context.Context, team_id int32, subscription_id int32) (e0 error) {
	request := DeleteTeamSubscriptionRequest{
		SubscriptionId: subscription_id,
		TeamId:         team_id,
	}
	response, err := e.DeleteTeamSubscriptionEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteTeamSubscriptionResponse).E0
}

// GetTeam implements Service. Primarily useful in a client.
func (e Endpoints) GetTeam(ctx context.Context, team_id int32) (e0 error, m1 proto.Team) {
	request := GetTeamRequest{TeamId: team_id}
	response, err := e.GetTeamEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetTeamResponse).E0, response.(GetTeamResponse).M1
}

// GetTeamProfile implements Service. Primarily useful in a client.
func (e Endpoints) GetTeamProfile(ctx context.Context, team_id int32) (e0 error, m1 proto.Profile) {
	request := GetTeamProfileRequest{TeamId: team_id}
	response, err := e.GetTeamProfileEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetTeamProfileResponse).E0, response.(GetTeamProfileResponse).M1
}

// GetTeamSubscriptions implements Service. Primarily useful in a client.
func (e Endpoints) GetTeamSubscriptions(ctx context.Context, team_id int32) (e0 error, m1 []model.Subscriptions) {
	request := GetTeamSubscriptionsRequest{TeamId: team_id}
	response, err := e.GetTeamSubscriptionsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetTeamSubscriptionsResponse).E0, response.(GetTeamSubscriptionsResponse).M1
}

// GetTeamMembers implements Service. Primarily useful in a client.
func (e Endpoints) GetTeamMembers(ctx context.Context, team_id int32) (e0 error, m1 []model.User) {
	request := GetTeamMembersRequest{TeamId: team_id}
	response, err := e.GetTeamMembersEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetTeamMembersResponse).E0, response.(GetTeamMembersResponse).M1
}

// GetTeamAdvisors implements Service. Primarily useful in a client.
func (e Endpoints) GetTeamAdvisors(ctx context.Context, team_id int32) (e0 error, m1 []model.User) {
	request := GetTeamAdvisorsRequest{TeamId: team_id}
	response, err := e.GetTeamAdvisorsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetTeamAdvisorsResponse).E0, response.(GetTeamAdvisorsResponse).M1
}

// GetTeamAdmin implements Service. Primarily useful in a client.
func (e Endpoints) GetTeamAdmin(ctx context.Context, team_id int32) (e0 error, m1 proto.User) {
	request := GetTeamAdminRequest{TeamId: team_id}
	response, err := e.GetTeamAdminEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetTeamAdminResponse).E0, response.(GetTeamAdminResponse).M1
}

// GetTeams implements Service. Primarily useful in a client.
func (e Endpoints) GetTeams(ctx context.Context, limit int32) (e0 error, m1 []model.Team) {
	request := GetTeamsRequest{Limit: limit}
	response, err := e.GetTeamsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetTeamsResponse).E0, response.(GetTeamsResponse).M1
}

// GetTeamsByType implements Service. Primarily useful in a client.
func (e Endpoints) GetTeamsByType(ctx context.Context, team_type string, limit int32) (e0 error, m1 []model.Team) {
	request := GetTeamsByTypeRequest{
		Limit:    limit,
		TeamType: team_type,
	}
	response, err := e.GetTeamsByTypeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetTeamsByTypeResponse).E0, response.(GetTeamsByTypeResponse).M1
}

// GetTeamsByIndustry implements Service. Primarily useful in a client.
func (e Endpoints) GetTeamsByIndustry(ctx context.Context, industry string, limit int32) (e0 error, m1 []model.Team) {
	request := GetTeamsByIndustryRequest{
		Industry: industry,
		Limit:    limit,
	}
	response, err := e.GetTeamsByIndustryEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetTeamsByIndustryResponse).E0, response.(GetTeamsByIndustryResponse).M1
}

// GetTeamsByNumberOfEmployees implements Service. Primarily useful in a client.
func (e Endpoints) GetTeamsByNumberOfEmployees(ctx context.Context, numEmployees int32, limit int32) (e0 error, m1 []model.Team) {
	request := GetTeamsByNumberOfEmployeesRequest{
		Limit:        limit,
		NumEmployees: numEmployees,
	}
	response, err := e.GetTeamsByNumberOfEmployeesEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetTeamsByNumberOfEmployeesResponse).E0, response.(GetTeamsByNumberOfEmployeesResponse).M1
}

// GetTeamProfiles implements Service. Primarily useful in a client.
func (e Endpoints) GetTeamProfiles(ctx context.Context, limit int32) (e0 error, m1 []model.TeamProfile) {
	request := GetTeamProfilesRequest{Limit: limit}
	response, err := e.GetTeamProfilesEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetTeamProfilesResponse).E0, response.(GetTeamProfilesResponse).M1
}

// GetTeamsByTags implements Service. Primarily useful in a client.
func (e Endpoints) GetTeamsByTags(ctx context.Context, tags string, limit int32) (e0 error, m1 []model.TeamProfile) {
	request := GetTeamsByTagsRequest{
		Limit: limit,
		Tags:  tags,
	}
	response, err := e.GetTeamsByTagsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetTeamsByTagsResponse).E0, response.(GetTeamsByTagsResponse).M1
}

// CreateGroup implements Service. Primarily useful in a client.
func (e Endpoints) CreateGroup(ctx context.Context, group proto.Group, admin proto.User) (e0 error) {
	request := CreateGroupRequest{
		Admin: admin,
		Group: group,
	}
	response, err := e.CreateGroupEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateGroupResponse).E0
}

// AddGroupMember implements Service. Primarily useful in a client.
func (e Endpoints) AddGroupMember(ctx context.Context, group_id int32, member proto.User) (e0 error) {
	request := AddGroupMemberRequest{
		GroupId: group_id,
		Member:  member,
	}
	response, err := e.AddGroupMemberEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddGroupMemberResponse).E0
}

// UpdateGroup implements Service. Primarily useful in a client.
func (e Endpoints) UpdateGroup(ctx context.Context, group_id int32, group proto.Group) (e0 error) {
	request := UpdateGroupRequest{
		Group:   group,
		GroupId: group_id,
	}
	response, err := e.UpdateGroupEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateGroupResponse).E0
}

// UpdateGroupMember implements Service. Primarily useful in a client.
func (e Endpoints) UpdateGroupMember(ctx context.Context, group_id int32, group_member_id int32, member proto.User) (e0 error) {
	request := UpdateGroupMemberRequest{
		GroupId:       group_id,
		GroupMemberId: group_member_id,
		Member:        member,
	}
	response, err := e.UpdateGroupMemberEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateGroupMemberResponse).E0
}

// DeleteGroup implements Service. Primarily useful in a client.
func (e Endpoints) DeleteGroup(ctx context.Context, group_id int32) (e0 error) {
	request := DeleteGroupRequest{GroupId: group_id}
	response, err := e.DeleteGroupEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteGroupResponse).E0
}

// DeleteGroupMember implements Service. Primarily useful in a client.
func (e Endpoints) DeleteGroupMember(ctx context.Context, group_id int32, member_id int32) (e0 error) {
	request := DeleteGroupMemberRequest{
		GroupId:  group_id,
		MemberId: member_id,
	}
	response, err := e.DeleteGroupMemberEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteGroupMemberResponse).E0
}

// GetGroup implements Service. Primarily useful in a client.
func (e Endpoints) GetGroup(ctx context.Context, group_id int32) (e0 error, m1 proto.Group) {
	request := GetGroupRequest{GroupId: group_id}
	response, err := e.GetGroupEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetGroupResponse).E0, response.(GetGroupResponse).M1
}

// GetGroups implements Service. Primarily useful in a client.
func (e Endpoints) GetGroups(ctx context.Context, limit int32) (e0 error, m1 []model.Group) {
	request := GetGroupsRequest{Limit: limit}
	response, err := e.GetGroupsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetGroupsResponse).E0, response.(GetGroupsResponse).M1
}

// GetGroupsByType implements Service. Primarily useful in a client.
func (e Endpoints) GetGroupsByType(ctx context.Context, group_type string, limit int32) (e0 error, m1 []model.Group) {
	request := GetGroupsByTypeRequest{
		GroupType: group_type,
		Limit:     limit,
	}
	response, err := e.GetGroupsByTypeEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetGroupsByTypeResponse).E0, response.(GetGroupsByTypeResponse).M1
}

// GetGroupsByNumMembers implements Service. Primarily useful in a client.
func (e Endpoints) GetGroupsByNumMembers(ctx context.Context, maxNumEmployees int32, limit int32) (e0 error, m1 []model.Group) {
	request := GetGroupsByNumMembersRequest{
		Limit:           limit,
		MaxNumEmployees: maxNumEmployees,
	}
	response, err := e.GetGroupsByNumMembersEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetGroupsByNumMembersResponse).E0, response.(GetGroupsByNumMembersResponse).M1
}

// GetGroupsByTags implements Service. Primarily useful in a client.
func (e Endpoints) GetGroupsByTags(ctx context.Context, tags string, limit int32) (e0 error, m1 []model.Group) {
	request := GetGroupsByTagsRequest{
		Limit: limit,
		Tags:  tags,
	}
	response, err := e.GetGroupsByTagsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetGroupsByTagsResponse).E0, response.(GetGroupsByTagsResponse).M1
}
