// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	endpoint "github.com/LensPlatform/micro/pkg/endpoint"
	pb "github.com/LensPlatform/micro/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	createUser                   grpc.Handler
	createProfile                grpc.Handler
	createSubscription           grpc.Handler
	createSubscriptions          grpc.Handler
	updateUser                   grpc.Handler
	updateUserProfile            grpc.Handler
	updateUserSubscription       grpc.Handler
	deleteUser                   grpc.Handler
	deleteUserProfile            grpc.Handler
	deleteSubscription           grpc.Handler
	getUser                      grpc.Handler
	getUsers                     grpc.Handler
	getUsersByAccountType        grpc.Handler
	getUsersByIntent             grpc.Handler
	getUserProfile               grpc.Handler
	getUserProfiles              grpc.Handler
	getUserProfilesByType        grpc.Handler
	getUserProfilesByNationality grpc.Handler
	getUserSubscriptions         grpc.Handler
	createTeam                   grpc.Handler
	creatTeamProfile             grpc.Handler
	createTeamSubscription       grpc.Handler
	addMemberToTeam              grpc.Handler
	addAdvisorToTeam             grpc.Handler
	updateTeam                   grpc.Handler
	updateTeamAdmin              grpc.Handler
	updateTeamMember             grpc.Handler
	updateTeamAdvisor            grpc.Handler
	updateTeamProfile            grpc.Handler
	updateTeamSubscription       grpc.Handler
	deleteTeam                   grpc.Handler
	deleteTeamProfile            grpc.Handler
	deleteTeamMember             grpc.Handler
	deleteTeamAdvisor            grpc.Handler
	deleteTeamAdmin              grpc.Handler
	deleteTeamSubscription       grpc.Handler
	getTeam                      grpc.Handler
	getTeamProfile               grpc.Handler
	getTeamSubscriptions         grpc.Handler
	getTeamMembers               grpc.Handler
	getTeamAdvisors              grpc.Handler
	getTeamAdmin                 grpc.Handler
	getTeams                     grpc.Handler
	getTeamsByType               grpc.Handler
	getTeamsByIndustry           grpc.Handler
	getTeamsByNumberOfEmployees  grpc.Handler
	getTeamProfiles              grpc.Handler
	getTeamsByTags               grpc.Handler
	createGroup                  grpc.Handler
	addGroupMember               grpc.Handler
	updateGroup                  grpc.Handler
	updateGroupMember            grpc.Handler
	deleteGroup                  grpc.Handler
	deleteGroupMember            grpc.Handler
	getGroup                     grpc.Handler
	getGroups                    grpc.Handler
	getGroupsByType              grpc.Handler
	getGroupsByNumMembers        grpc.Handler
	getGroupsByTags              grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.MicroServer {
	return &grpcServer{
		addAdvisorToTeam:             makeAddAdvisorToTeamHandler(endpoints, options["AddAdvisorToTeam"]),
		addGroupMember:               makeAddGroupMemberHandler(endpoints, options["AddGroupMember"]),
		addMemberToTeam:              makeAddMemberToTeamHandler(endpoints, options["AddMemberToTeam"]),
		creatTeamProfile:             makeCreatTeamProfileHandler(endpoints, options["CreatTeamProfile"]),
		createGroup:                  makeCreateGroupHandler(endpoints, options["CreateGroup"]),
		createProfile:                makeCreateProfileHandler(endpoints, options["CreateProfile"]),
		createSubscription:           makeCreateSubscriptionHandler(endpoints, options["CreateSubscription"]),
		createSubscriptions:          makeCreateSubscriptionsHandler(endpoints, options["CreateSubscriptions"]),
		createTeam:                   makeCreateTeamHandler(endpoints, options["CreateTeam"]),
		createTeamSubscription:       makeCreateTeamSubscriptionHandler(endpoints, options["CreateTeamSubscription"]),
		createUser:                   makeCreateUserHandler(endpoints, options["CreateUser"]),
		deleteGroup:                  makeDeleteGroupHandler(endpoints, options["DeleteGroup"]),
		deleteGroupMember:            makeDeleteGroupMemberHandler(endpoints, options["DeleteGroupMember"]),
		deleteSubscription:           makeDeleteSubscriptionHandler(endpoints, options["DeleteSubscription"]),
		deleteTeam:                   makeDeleteTeamHandler(endpoints, options["DeleteTeam"]),
		deleteTeamAdmin:              makeDeleteTeamAdminHandler(endpoints, options["DeleteTeamAdmin"]),
		deleteTeamAdvisor:            makeDeleteTeamAdvisorHandler(endpoints, options["DeleteTeamAdvisor"]),
		deleteTeamMember:             makeDeleteTeamMemberHandler(endpoints, options["DeleteTeamMember"]),
		deleteTeamProfile:            makeDeleteTeamProfileHandler(endpoints, options["DeleteTeamProfile"]),
		deleteTeamSubscription:       makeDeleteTeamSubscriptionHandler(endpoints, options["DeleteTeamSubscription"]),
		deleteUser:                   makeDeleteUserHandler(endpoints, options["DeleteUser"]),
		deleteUserProfile:            makeDeleteUserProfileHandler(endpoints, options["DeleteUserProfile"]),
		getGroup:                     makeGetGroupHandler(endpoints, options["GetGroup"]),
		getGroups:                    makeGetGroupsHandler(endpoints, options["GetGroups"]),
		getGroupsByNumMembers:        makeGetGroupsByNumMembersHandler(endpoints, options["GetGroupsByNumMembers"]),
		getGroupsByTags:              makeGetGroupsByTagsHandler(endpoints, options["GetGroupsByTags"]),
		getGroupsByType:              makeGetGroupsByTypeHandler(endpoints, options["GetGroupsByType"]),
		getTeam:                      makeGetTeamHandler(endpoints, options["GetTeam"]),
		getTeamAdmin:                 makeGetTeamAdminHandler(endpoints, options["GetTeamAdmin"]),
		getTeamAdvisors:              makeGetTeamAdvisorsHandler(endpoints, options["GetTeamAdvisors"]),
		getTeamMembers:               makeGetTeamMembersHandler(endpoints, options["GetTeamMembers"]),
		getTeamProfile:               makeGetTeamProfileHandler(endpoints, options["GetTeamProfile"]),
		getTeamProfiles:              makeGetTeamProfilesHandler(endpoints, options["GetTeamProfiles"]),
		getTeamSubscriptions:         makeGetTeamSubscriptionsHandler(endpoints, options["GetTeamSubscriptions"]),
		getTeams:                     makeGetTeamsHandler(endpoints, options["GetTeams"]),
		getTeamsByIndustry:           makeGetTeamsByIndustryHandler(endpoints, options["GetTeamsByIndustry"]),
		getTeamsByNumberOfEmployees:  makeGetTeamsByNumberOfEmployeesHandler(endpoints, options["GetTeamsByNumberOfEmployees"]),
		getTeamsByTags:               makeGetTeamsByTagsHandler(endpoints, options["GetTeamsByTags"]),
		getTeamsByType:               makeGetTeamsByTypeHandler(endpoints, options["GetTeamsByType"]),
		getUser:                      makeGetUserHandler(endpoints, options["GetUser"]),
		getUserProfile:               makeGetUserProfileHandler(endpoints, options["GetUserProfile"]),
		getUserProfiles:              makeGetUserProfilesHandler(endpoints, options["GetUserProfiles"]),
		getUserProfilesByNationality: makeGetUserProfilesByNationalityHandler(endpoints, options["GetUserProfilesByNationality"]),
		getUserProfilesByType:        makeGetUserProfilesByTypeHandler(endpoints, options["GetUserProfilesByType"]),
		getUserSubscriptions:         makeGetUserSubscriptionsHandler(endpoints, options["GetUserSubscriptions"]),
		getUsers:                     makeGetUsersHandler(endpoints, options["GetUsers"]),
		getUsersByAccountType:        makeGetUsersByAccountTypeHandler(endpoints, options["GetUsersByAccountType"]),
		getUsersByIntent:             makeGetUsersByIntentHandler(endpoints, options["GetUsersByIntent"]),
		updateGroup:                  makeUpdateGroupHandler(endpoints, options["UpdateGroup"]),
		updateGroupMember:            makeUpdateGroupMemberHandler(endpoints, options["UpdateGroupMember"]),
		updateTeam:                   makeUpdateTeamHandler(endpoints, options["UpdateTeam"]),
		updateTeamAdmin:              makeUpdateTeamAdminHandler(endpoints, options["UpdateTeamAdmin"]),
		updateTeamAdvisor:            makeUpdateTeamAdvisorHandler(endpoints, options["UpdateTeamAdvisor"]),
		updateTeamMember:             makeUpdateTeamMemberHandler(endpoints, options["UpdateTeamMember"]),
		updateTeamProfile:            makeUpdateTeamProfileHandler(endpoints, options["UpdateTeamProfile"]),
		updateTeamSubscription:       makeUpdateTeamSubscriptionHandler(endpoints, options["UpdateTeamSubscription"]),
		updateUser:                   makeUpdateUserHandler(endpoints, options["UpdateUser"]),
		updateUserProfile:            makeUpdateUserProfileHandler(endpoints, options["UpdateUserProfile"]),
		updateUserSubscription:       makeUpdateUserSubscriptionHandler(endpoints, options["UpdateUserSubscription"]),
	}
}