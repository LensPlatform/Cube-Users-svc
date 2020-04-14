// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	endpoint "github.com/LensPlatform/micro/pkg/endpoint"
	http "github.com/go-kit/kit/transport/http"
	http1 "net/http"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := http1.NewServeMux()
	makeCreateUserHandler(m, endpoints, options["CreateUser"])
	makeCreateProfileHandler(m, endpoints, options["CreateProfile"])
	makeCreateSubscriptionHandler(m, endpoints, options["CreateSubscription"])
	makeCreateSubscriptionsHandler(m, endpoints, options["CreateSubscriptions"])
	makeUpdateUserHandler(m, endpoints, options["UpdateUser"])
	makeUpdateUserProfileHandler(m, endpoints, options["UpdateUserProfile"])
	makeUpdateUserSubscriptionHandler(m, endpoints, options["UpdateUserSubscription"])
	makeDeleteUserHandler(m, endpoints, options["DeleteUser"])
	makeDeleteUserProfileHandler(m, endpoints, options["DeleteUserProfile"])
	makeDeleteSubscriptionHandler(m, endpoints, options["DeleteSubscription"])
	makeGetUserHandler(m, endpoints, options["GetUser"])
	makeGetUsersHandler(m, endpoints, options["GetUsers"])
	makeGetUsersByAccountTypeHandler(m, endpoints, options["GetUsersByAccountType"])
	makeGetUsersByIntentHandler(m, endpoints, options["GetUsersByIntent"])
	makeGetUserProfileHandler(m, endpoints, options["GetUserProfile"])
	makeGetUserProfilesHandler(m, endpoints, options["GetUserProfiles"])
	makeGetUserProfilesByTypeHandler(m, endpoints, options["GetUserProfilesByType"])
	makeGetUserProfilesByNationalityHandler(m, endpoints, options["GetUserProfilesByNationality"])
	makeGetUserSubscriptionsHandler(m, endpoints, options["GetUserSubscriptions"])
	makeCreateTeamHandler(m, endpoints, options["CreateTeam"])
	makeCreatTeamProfileHandler(m, endpoints, options["CreatTeamProfile"])
	makeCreateTeamSubscriptionHandler(m, endpoints, options["CreateTeamSubscription"])
	makeAddMemberToTeamHandler(m, endpoints, options["AddMemberToTeam"])
	makeAddAdvisorToTeamHandler(m, endpoints, options["AddAdvisorToTeam"])
	makeUpdateTeamHandler(m, endpoints, options["UpdateTeam"])
	makeUpdateTeamAdminHandler(m, endpoints, options["UpdateTeamAdmin"])
	makeUpdateTeamMemberHandler(m, endpoints, options["UpdateTeamMember"])
	makeUpdateTeamAdvisorHandler(m, endpoints, options["UpdateTeamAdvisor"])
	makeUpdateTeamProfileHandler(m, endpoints, options["UpdateTeamProfile"])
	makeUpdateTeamSubscriptionHandler(m, endpoints, options["UpdateTeamSubscription"])
	makeDeleteTeamHandler(m, endpoints, options["DeleteTeam"])
	makeDeleteTeamProfileHandler(m, endpoints, options["DeleteTeamProfile"])
	makeDeleteTeamMemberHandler(m, endpoints, options["DeleteTeamMember"])
	makeDeleteTeamAdvisorHandler(m, endpoints, options["DeleteTeamAdvisor"])
	makeDeleteTeamAdminHandler(m, endpoints, options["DeleteTeamAdmin"])
	makeDeleteTeamSubscriptionHandler(m, endpoints, options["DeleteTeamSubscription"])
	makeGetTeamHandler(m, endpoints, options["GetTeam"])
	makeGetTeamProfileHandler(m, endpoints, options["GetTeamProfile"])
	makeGetTeamSubscriptionsHandler(m, endpoints, options["GetTeamSubscriptions"])
	makeGetTeamMembersHandler(m, endpoints, options["GetTeamMembers"])
	makeGetTeamAdvisorsHandler(m, endpoints, options["GetTeamAdvisors"])
	makeGetTeamAdminHandler(m, endpoints, options["GetTeamAdmin"])
	makeGetTeamsHandler(m, endpoints, options["GetTeams"])
	makeGetTeamsByTypeHandler(m, endpoints, options["GetTeamsByType"])
	makeGetTeamsByIndustryHandler(m, endpoints, options["GetTeamsByIndustry"])
	makeGetTeamsByNumberOfEmployeesHandler(m, endpoints, options["GetTeamsByNumberOfEmployees"])
	makeGetTeamProfilesHandler(m, endpoints, options["GetTeamProfiles"])
	makeGetTeamsByTagsHandler(m, endpoints, options["GetTeamsByTags"])
	makeCreateGroupHandler(m, endpoints, options["CreateGroup"])
	makeAddGroupMemberHandler(m, endpoints, options["AddGroupMember"])
	makeUpdateGroupHandler(m, endpoints, options["UpdateGroup"])
	makeUpdateGroupMemberHandler(m, endpoints, options["UpdateGroupMember"])
	makeDeleteGroupHandler(m, endpoints, options["DeleteGroup"])
	makeDeleteGroupMemberHandler(m, endpoints, options["DeleteGroupMember"])
	makeGetGroupHandler(m, endpoints, options["GetGroup"])
	makeGetGroupsHandler(m, endpoints, options["GetGroups"])
	makeGetGroupsByTypeHandler(m, endpoints, options["GetGroupsByType"])
	makeGetGroupsByNumMembersHandler(m, endpoints, options["GetGroupsByNumMembers"])
	makeGetGroupsByTagsHandler(m, endpoints, options["GetGroupsByTags"])
	return m
}
