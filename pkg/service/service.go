package service

import (
	"context"
	"errors"
	"os"

	"github.com/LensPlatform/micro/pkg/database"
	"github.com/LensPlatform/micro/pkg/helper"
	model "github.com/LensPlatform/micro/pkg/models/proto"
	"github.com/go-kit/kit/log"
)

// MicroService describes the service.
type MicroService interface {
	CreateUser(ctx context.Context, user model.User) error
	CreateProfile(ctx context.Context, user_id int32, profile model.Profile) error
	CreateSubscription(ctx context.Context, user_id int32, subscription model.Subscriptions) error
	CreateSubscriptions(ctx context.Context, user_id, subscriptions []model.Subscriptions) error
	UpdateUser(ctx context.Context, user_id int32, user model.User) (error, model.User)
	UpdateUserProfile(ctx context.Context, user_id int32, profile_id int32, profile model.Profile) (error, model.Profile)
	UpdateUserSubscription(ctx context.Context, user_id int32, subscription_id int32, subscription model.Subscriptions) (error, model.Subscriptions)
	DeleteUser(ctx context.Context, user_id int32) error
	DeleteUserProfile(ctx context.Context, user_id, profile_id int32) error
	DeleteSubscription(ctx context.Context, user_id, subscription_id int32) error

	GetUser(ctx context.Context, user_id int32) (error, model.User)
	GetUsers(ctx context.Context, limit int32) (error, []model.User)
	GetUsersByAccountType(ctx context.Context, accounType string, limit int32) (error, []model.User)
	GetUsersByIntent(ctx context.Context, intent string, limit int32) (error, []model.User)
	GetUserProfile(ctx context.Context, user_id int32) (error, model.Profile)
	GetUserProfiles(ctx context.Context, limit int32) (error, []model.Profile)
	GetUserProfilesByType(ctx context.Context, profileType string, limit int32) (error, []model.Profile)
	GetUserProfilesByNationality(ctx context.Context, nationality string, limit int32) (error, []model.Profile)
	GetUserSubscriptions(ctx context.Context, user_id int32) (error, []model.Subscriptions)

	CreateTeam(ctx context.Context, team model.Team, admin model.User) error
	CreatTeamProfile(ctx context.Context, team_id int32, profile model.TeamProfile) error
	CreateTeamSubscription(ctx context.Context, team_id int32, subscription model.Subscriptions) error
	AddMemberToTeam(ctx context.Context, team_id int32, member model.User) error

	AddAdvisorToTeam(ctx context.Context, team_id int32, advisor_id int32, advisor model.User) error
	UpdateTeam(ctx context.Context, team_id int32, team model.Team) (error, model.Team)
	UpdateTeamAdmin(ctx context.Context, team_id int32, admin_id int32, admin model.User) (error, model.Team)
	UpdateTeamMember(ctx context.Context, team_id int32, member_id int32, member model.User) (error, model.Team)
	UpdateTeamAdvisor(ctx context.Context, team_id int32, advisor_id int32, advisor model.User) (error, model.Team)
	UpdateTeamProfile(ctx context.Context, team_id int32, profile_id int32, profile model.TeamProfile) (error, model.Team)
	UpdateTeamSubscription(ctx context.Context, team_id int32, subscription_id int32, subcription model.Subscriptions) (error, model.Team)

	DeleteTeam(ctx context.Context, team_id int32) error
	DeleteTeamProfile(ctx context.Context, team_id int32, profile_id int32) error
	DeleteTeamMember(ctx context.Context, team_id int32, team_member_id int32) error
	DeleteTeamAdvisor(ctx context.Context, team_id, advisor_id int32) error
	DeleteTeamAdmin(ctx context.Context, team_id, admin_id int32) error
	DeleteTeamSubscription(ctx context.Context, team_id, subscription_id int32) error

	GetTeam(ctx context.Context, team_id int32) (error, model.Team)
	GetTeamProfile(ctx context.Context, team_id int32) (error, model.Profile)
	GetTeamSubscriptions(ctx context.Context, team_id int32) (error, []model.Subscriptions)
	GetTeamMembers(ctx context.Context, team_id int32) (error, []model.User)
	GetTeamAdvisors(ctx context.Context, team_id int32) (error, []model.User)
	GetTeamAdmin(ctx context.Context, team_id int32) (error, model.User)
	GetTeams(ctx context.Context, limit int32) (error, []model.Team)
	GetTeamsByType(ctx context.Context, team_type string, limit int32) (error, []model.Team)
	GetTeamsByIndustry(ctx context.Context, industry string, limit int32) (error, []model.Team)
	GetTeamsByNumberOfEmployees(ctx context.Context, numEmployees int32, limit int32) (error, []model.Team)
	GetTeamProfiles(ctx context.Context, limit int32) (error, []model.TeamProfile)
	GetTeamsByTags(ctx context.Context, tags string, limit int32) (error, []model.TeamProfile)

	CreateGroup(ctx context.Context, group model.Group, admin model.User) error
	AddGroupMember(ctx context.Context, group_id int32, member model.User) error
	UpdateGroup(ctx context.Context, group_id int32, group model.Group) error
	UpdateGroupMember(ctx context.Context, group_id int32, group_member_id int32, member model.User) error
	DeleteGroup(ctx context.Context, group_id int32) error
	DeleteGroupMember(ctx context.Context, group_id, member_id int32) error
	GetGroup(ctx context.Context, group_id int32) (error, model.Group)
	GetGroups(ctx context.Context, limit int32) (error, []model.Group)
	GetGroupsByType(ctx context.Context, group_type string, limit int32) (error, []model.Group)
	GetGroupsByNumMembers(ctx context.Context, maxNumEmployees, limit int32) (error, []model.Group)
	GetGroupsByTags(ctx context.Context, tags string, limit int32) (error, []model.Group)
}

type basicMicroService struct {
	db *Database
}

func (b *basicMicroService) CreateUser(ctx context.Context, user model.User) (e0 error) {
	if user.UserName == "" || user.Email == "" {
		return errors.New("Invalid input parameters. User email and username does not exist.")
	}

	user, e0 = helper.ValidateAndHashPassword(user)
	if e0 != nil {
		return e0
	}

	e0 = b.db.CreateUser(user)
	if e0 != nil {
		return e0
	}
	return nil
}
func (b *basicMicroService) CreateProfile(ctx context.Context, user_id int32, profile model.Profile) (e0 error) {
	_, exists := b.db.DoesUserExists(user_id, "", "")
	if exists {
		// query the database for the profile and check for existence
		_, profileExists := b.db.DoesUserProfileExists(user_id)

		if profileExists {
			return errors.New("User profile already exists")
		} else {
			return b.db.CreateUserProfile(user_id, profile)
		}
	} else {
		return errors.New("User account does not exist. Please create a user account before creating a profile")
	}

	return nil
}
func (b *basicMicroService) CreateSubscription(ctx context.Context, user_id int32, subscription model.Subscriptions) (e0 error) {
	// TODO implement the business logic of CreateSubscription
	return e0
}
func (b *basicMicroService) CreateSubscriptions(ctx context.Context, user_id []model.Subscriptions, subscriptions []model.Subscriptions) (e0 error) {
	// TODO implement the business logic of CreateSubscriptions
	return e0
}
func (b *basicMicroService) UpdateUser(ctx context.Context, user_id int32, user model.User) (e0 error, m1 model.User) {
	// TODO implement the business logic of UpdateUser
	return e0, m1
}
func (b *basicMicroService) UpdateUserProfile(ctx context.Context, user_id int32, profile_id int32, profile model.Profile) (e0 error, m1 model.Profile) {
	// TODO implement the business logic of UpdateUserProfile
	return e0, m1
}
func (b *basicMicroService) UpdateUserSubscription(ctx context.Context, user_id int32, subscription_id int32, subscription model.Subscriptions) (e0 error, m1 model.Subscriptions) {
	// TODO implement the business logic of UpdateUserSubscription
	return e0, m1
}
func (b *basicMicroService) DeleteUser(ctx context.Context, user_id int32) (e0 error) {
	// TODO implement the business logic of DeleteUser
	return e0
}
func (b *basicMicroService) DeleteUserProfile(ctx context.Context, user_id int32, profile_id int32) (e0 error) {
	// TODO implement the business logic of DeleteUserProfile
	return e0
}
func (b *basicMicroService) DeleteSubscription(ctx context.Context, user_id int32, subscription_id int32) (e0 error) {
	// TODO implement the business logic of DeleteSubscription
	return e0
}
func (b *basicMicroService) GetUser(ctx context.Context, user_id int32) (e0 error, m1 model.User) {
	e0, user := b.db.GetUserById(user_id)
	if e0 != nil {
		b.db.Logger.Log(e0.Error())
	}
	m1 = *user
	return e0, m1
}
func (b *basicMicroService) GetUsers(ctx context.Context, limit int32) (e0 error, m1 []model.User) {
	// TODO implement the business logic of GetUsers
	return e0, m1
}
func (b *basicMicroService) GetUsersByAccountType(ctx context.Context, accounType string, limit int32) (e0 error, m1 []model.User) {
	// TODO implement the business logic of GetUsersByAccountType
	return e0, m1
}
func (b *basicMicroService) GetUsersByIntent(ctx context.Context, intent string, limit int32) (e0 error, m1 []model.User) {
	// TODO implement the business logic of GetUsersByIntent
	return e0, m1
}
func (b *basicMicroService) GetUserProfile(ctx context.Context, user_id int32) (e0 error, m1 model.Profile) {
	// TODO implement the business logic of GetUserProfile
	return e0, m1
}
func (b *basicMicroService) GetUserProfiles(ctx context.Context, limit int32) (e0 error, m1 []model.Profile) {
	// TODO implement the business logic of GetUserProfiles
	return e0, m1
}
func (b *basicMicroService) GetUserProfilesByType(ctx context.Context, profileType string, limit int32) (e0 error, m1 []model.Profile) {
	// TODO implement the business logic of GetUserProfilesByType
	return e0, m1
}
func (b *basicMicroService) GetUserProfilesByNationality(ctx context.Context, nationality string, limit int32) (e0 error, m1 []model.Profile) {
	// TODO implement the business logic of GetUserProfilesByNationality
	return e0, m1
}
func (b *basicMicroService) GetUserSubscriptions(ctx context.Context, user_id int32) (e0 error, m1 []model.Subscriptions) {
	// TODO implement the business logic of GetUserSubscriptions
	return e0, m1
}
func (b *basicMicroService) CreateTeam(ctx context.Context, team model.Team, admin model.User) (e0 error) {
	// TODO implement the business logic of CreateTeam
	return e0
}
func (b *basicMicroService) CreatTeamProfile(ctx context.Context, team_id int32, profile model.TeamProfile) (e0 error) {
	// TODO implement the business logic of CreatTeamProfile
	return e0
}
func (b *basicMicroService) CreateTeamSubscription(ctx context.Context, team_id int32, subscription model.Subscriptions) (e0 error) {
	// TODO implement the business logic of CreateTeamSubscription
	return e0
}
func (b *basicMicroService) AddMemberToTeam(ctx context.Context, team_id int32, member model.User) (e0 error) {
	// TODO implement the business logic of AddMemberToTeam
	return e0
}
func (b *basicMicroService) AddAdvisorToTeam(ctx context.Context, team_id int32, advisor_id int32, advisor model.User) (e0 error) {
	// TODO implement the business logic of AddAdvisorToTeam
	return e0
}
func (b *basicMicroService) UpdateTeam(ctx context.Context, team_id int32, team model.Team) (e0 error, m1 model.Team) {
	// TODO implement the business logic of UpdateTeam
	return e0, m1
}
func (b *basicMicroService) UpdateTeamAdmin(ctx context.Context, team_id int32, admin_id int32, admin model.User) (e0 error, m1 model.Team) {
	// TODO implement the business logic of UpdateTeamAdmin
	return e0, m1
}
func (b *basicMicroService) UpdateTeamMember(ctx context.Context, team_id int32, member_id int32, member model.User) (e0 error, m1 model.Team) {
	// TODO implement the business logic of UpdateTeamMember
	return e0, m1
}
func (b *basicMicroService) UpdateTeamAdvisor(ctx context.Context, team_id int32, advisor_id int32, advisor model.User) (e0 error, m1 model.Team) {
	// TODO implement the business logic of UpdateTeamAdvisor
	return e0, m1
}
func (b *basicMicroService) UpdateTeamProfile(ctx context.Context, team_id int32, profile_id int32, profile model.TeamProfile) (e0 error, m1 model.Team) {
	// TODO implement the business logic of UpdateTeamProfile
	return e0, m1
}
func (b *basicMicroService) UpdateTeamSubscription(ctx context.Context, team_id int32, subscription_id int32, subcription model.Subscriptions) (e0 error, m1 model.Team) {
	// TODO implement the business logic of UpdateTeamSubscription
	return e0, m1
}
func (b *basicMicroService) DeleteTeam(ctx context.Context, team_id int32) (e0 error) {
	// TODO implement the business logic of DeleteTeam
	return e0
}
func (b *basicMicroService) DeleteTeamProfile(ctx context.Context, team_id int32, profile_id int32) (e0 error) {
	// TODO implement the business logic of DeleteTeamProfile
	return e0
}
func (b *basicMicroService) DeleteTeamMember(ctx context.Context, team_id int32, team_member_id int32) (e0 error) {
	// TODO implement the business logic of DeleteTeamMember
	return e0
}
func (b *basicMicroService) DeleteTeamAdvisor(ctx context.Context, team_id int32, advisor_id int32) (e0 error) {
	// TODO implement the business logic of DeleteTeamAdvisor
	return e0
}
func (b *basicMicroService) DeleteTeamAdmin(ctx context.Context, team_id int32, admin_id int32) (e0 error) {
	// TODO implement the business logic of DeleteTeamAdmin
	return e0
}
func (b *basicMicroService) DeleteTeamSubscription(ctx context.Context, team_id int32, subscription_id int32) (e0 error) {
	// TODO implement the business logic of DeleteTeamSubscription
	return e0
}
func (b *basicMicroService) GetTeam(ctx context.Context, team_id int32) (e0 error, m1 model.Team) {
	// TODO implement the business logic of GetTeam
	return e0, m1
}
func (b *basicMicroService) GetTeamProfile(ctx context.Context, team_id int32) (e0 error, m1 model.Profile) {
	// TODO implement the business logic of GetTeamProfile
	return e0, m1
}
func (b *basicMicroService) GetTeamSubscriptions(ctx context.Context, team_id int32) (e0 error, m1 []model.Subscriptions) {
	// TODO implement the business logic of GetTeamSubscriptions
	return e0, m1
}
func (b *basicMicroService) GetTeamMembers(ctx context.Context, team_id int32) (e0 error, m1 []model.User) {
	// TODO implement the business logic of GetTeamMembers
	return e0, m1
}
func (b *basicMicroService) GetTeamAdvisors(ctx context.Context, team_id int32) (e0 error, m1 []model.User) {
	// TODO implement the business logic of GetTeamAdvisors
	return e0, m1
}
func (b *basicMicroService) GetTeamAdmin(ctx context.Context, team_id int32) (e0 error, m1 model.User) {
	// TODO implement the business logic of GetTeamAdmin
	return e0, m1
}
func (b *basicMicroService) GetTeams(ctx context.Context, limit int32) (e0 error, m1 []model.Team) {
	// TODO implement the business logic of GetTeams
	return e0, m1
}
func (b *basicMicroService) GetTeamsByType(ctx context.Context, team_type string, limit int32) (e0 error, m1 []model.Team) {
	// TODO implement the business logic of GetTeamsByType
	return e0, m1
}
func (b *basicMicroService) GetTeamsByIndustry(ctx context.Context, industry string, limit int32) (e0 error, m1 []model.Team) {
	// TODO implement the business logic of GetTeamsByIndustry
	return e0, m1
}
func (b *basicMicroService) GetTeamsByNumberOfEmployees(ctx context.Context, numEmployees int32, limit int32) (e0 error, m1 []model.Team) {
	// TODO implement the business logic of GetTeamsByNumberOfEmployees
	return e0, m1
}
func (b *basicMicroService) GetTeamProfiles(ctx context.Context, limit int32) (e0 error, m1 []model.TeamProfile) {
	// TODO implement the business logic of GetTeamProfiles
	return e0, m1
}
func (b *basicMicroService) GetTeamsByTags(ctx context.Context, tags string, limit int32) (e0 error, m1 []model.TeamProfile) {
	// TODO implement the business logic of GetTeamsByTags
	return e0, m1
}
func (b *basicMicroService) CreateGroup(ctx context.Context, group model.Group, admin model.User) (e0 error) {
	// TODO implement the business logic of CreateGroup
	return e0
}
func (b *basicMicroService) AddGroupMember(ctx context.Context, group_id int32, member model.User) (e0 error) {
	// TODO implement the business logic of AddGroupMember
	return e0
}
func (b *basicMicroService) UpdateGroup(ctx context.Context, group_id int32, group model.Group) (e0 error) {
	// TODO implement the business logic of UpdateGroup
	return e0
}
func (b *basicMicroService) UpdateGroupMember(ctx context.Context, group_id int32, group_member_id int32, member model.User) (e0 error) {
	// TODO implement the business logic of UpdateGroupMember
	return e0
}
func (b *basicMicroService) DeleteGroup(ctx context.Context, group_id int32) (e0 error) {
	// TODO implement the business logic of DeleteGroup
	return e0
}
func (b *basicMicroService) DeleteGroupMember(ctx context.Context, group_id int32, member_id int32) (e0 error) {
	// TODO implement the business logic of DeleteGroupMember
	return e0
}
func (b *basicMicroService) GetGroup(ctx context.Context, group_id int32) (e0 error, m1 model.Group) {
	// TODO implement the business logic of GetGroup
	return e0, m1
}
func (b *basicMicroService) GetGroups(ctx context.Context, limit int32) (e0 error, m1 []model.Group) {
	// TODO implement the business logic of GetGroups
	return e0, m1
}
func (b *basicMicroService) GetGroupsByType(ctx context.Context, group_type string, limit int32) (e0 error, m1 []model.Group) {
	// TODO implement the business logic of GetGroupsByType
	return e0, m1
}
func (b *basicMicroService) GetGroupsByNumMembers(ctx context.Context, maxNumEmployees int32, limit int32) (e0 error, m1 []model.Group) {
	// TODO implement the business logic of GetGroupsByNumMembers
	return e0, m1
}
func (b *basicMicroService) GetGroupsByTags(ctx context.Context, tags string, limit int32) (e0 error, m1 []model.Group) {
	// TODO implement the business logic of GetGroupsByTags
	return e0, m1
}

// NewBasicMicroService returns a naive, stateless implementation of MicroService.
func NewBasicMicroService(conn string, logger log.Logger) MicroService {
	err, db := New(conn, logger)
	if err != nil {
		logger.Log(err.Error())
		os.Exit(1)
	}
	return &basicMicroService{
		db: db,
	}
}

// New returns a MicroService with all of the expected middleware wired in.
func New(middleware []Middleware, conn string, logger log.Logger) MicroService {
	var svc MicroService = NewBasicMicroService(conn, logger)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
