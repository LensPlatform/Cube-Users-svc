package database

import (
	"os"

	"gopkg.in/gormigrate.v1"

	model "github.com/LensPlatform/micro/pkg/models/proto"
	log "github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
)

type IDatabase interface {
	CreateUser(user model.User) error

	CreateUserProfile(user_id int32, profile model.Profile) error
	GetUserById(user_id int32) (error, *model.User)
	DoesUserExists(user_id int32, username, email string) (error, bool)
	DoesUserProfileExists(user_id, profile_id int32) (error, bool)
	/*
		CreateUserSubscription(user_id int32, subscription model.Subscriptions) error
		CreateUserSubscriptions(user_id int32, subscriptions []model.Subscriptions) error

		UpdateUser(user_id int32, user model.User) (error, *model.User)
		UpdateUserProfile(user_id, profile_id int32, profile model.Profile) (error, *model.Profile)
		UpdateUserSubscription(user_id int32, subscription_id int32, subscription model.Subscriptions) (error, *model.Subscriptions)

		DeleteUser(user_id int32) error
		DeleteUserProfile(user_id, profile_id int32) error
		DeleteUserSubscription(user_id, subscription_id int32) error

		GetUserByUsername(username string) (error, *model.User)
		GetUserByEmail(email string) (error, *model.User)
		GetAllUsers(limit int32) (error, []*model.User)
		GetAllUsersByAccountType(account_type string, limit int32) (error, []*model.User)
		GetAllUsersByIntent(intent string, limit int32) (error, []*model.User)

		GetUserProfile(user_id, profile_id int32) (error, []*model.Profile)
		GetAllUserProfilesByType(profile_type string, limit int32) (error, []*model.Profile)
		GetAllUserProfilesByNationality(nationality string, limit int32) (error, []*model.Profile)
		GetUserSubscriptions(user_id, subscription_id int32) (error, []*model.Subscriptions)

		DoesUserSubscriptionExists(user_id, subscription_id int32) (error, bool)

		CreateTeam(team model.Team) error
		CreateTeamProfile(team_id int32, profile model.TeamProfile) error
		CreateTeamSubscription(team_id int32, subscription model.Subscriptions) error
		AddMemberToTeam(team_id int32, member_id int32, member model.User) error
		AddAdvisorToTeam(team_id int32, advisor_id int32, advisor model.User) error

		UpdateTeam(team_id int32, team model.Team)
		UpdateTeamProfile(team_id, team_profile_id int32, profile model.TeamProfile) error
		UpdateTeamAdmin(team_id, admin_id int32, admin model.User) error
		UpdateTeamMember(team_id, member_id int32, member model.User) error
		UpdateTeamAdvisor(team_id, advisor_id int32, advisor model.User) error
		UpdateTeamSubscription(team_id, subscription_id int32, subscription model.Subscriptions) error

		DeleteTeam(team_id int32) error
		DeleteTeamProfile(team_id, team_profile_id int32) error
		DeleteTeamMember(team_id, member_id int32) error
		DeleteTeamAdvisor(team_id, advisor_id int32) error
		DeleteTeamAdmin(team_id, admin_id int32) error

		GetTeam(team_id int32) (error, *model.Team)
		GetAllTeams(limit int32) (error, []*model.Team)
		GetAllTeamsByType(team_type string, limit int32) (error, []*model.Team)
		GetAllTeamsByIndustry(industry string, limit int32) (error, []*model.Team)
		GetAllTeamsByNumberOfEmployees(numEmployees, limit int32) (error, []*model.Team)

		GetTeamProfile(team_id, profile_id int32) (error, *model.TeamProfile)
		GetAllTeamProfiles(profiles int32) (error, []*model.TeamProfile)
		GetAllProfilesByTag(tag string, limit int32) (error, []*model.TeamProfile)

		GetTeamSubscription(team_id, subscription_id int32) (error, *model.Subscriptions)
		GetTeamSubscriptions(team_id int32) (error, []*model.Subscriptions)
	*/
}

type Database struct {
	Engine *gorm.DB
	Logger log.Logger
}

var (
	postgres = "postgres"
)

// New creates a database connection and returns the conenction object
func New(conn string, logger log.Logger) (error, *Database) {
	db, err := gorm.Open(postgres, conn)
	if err != nil {
		logger.Log(err.Error())
		os.Exit(1)
	}

	logger.Log("Successfully connected to the database")

	db.SingularTable(true)
	db.LogMode(false)
	db = db.Set("gorm:auto_preload", true)

	logger.Log("Auto Migrating database models")

	err = CreateTablesOrMigrateSchemas(db, logger)
	if err != nil {
		logger.Log(err.Error())
		os.Exit(1)
	}

	logger.Log("Auto Migration of database models complete")

	return nil, &Database{
		Engine: db,
		Logger: logger,
	}
}

// CreateTablesOrMigrateSchemas creates a given set of models based on a schema
// if it does not exist or migrates the model schemas to the latest version
func CreateTablesOrMigrateSchemas(db *gorm.DB, logger log.Logger) error {
	migration := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20200416",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(model.AddressORM{}, model.EducationORM{}, model.MediaORM{}, model.SubscriptionsORM{}, model.SocialMediaORM{},
					model.DetailsORM{}, model.ExperienceORM{}, model.InvestmentORM{}, model.UserORM{}, model.ProfileORM{}, model.GroupORM{},
					model.TeamORM{}, model.TeamProfileORM{}, model.InvestorDetailORM{}, model.StartupDetailORM{}, model.SettingsORM{}, model.LoginActivityORM{},
					model.PaymentsORM{}, model.CardORM{}, model.PinORM{}, model.Privacy{}, model.NotificationORM{}, model.PostAndCommentsPushNotificationORM{},
					model.FollowingAndFollowersPushNotificationORM{}, model.DirectMessagesPushNotificationORM{}, model.EmailAndSmsPushNotificationORM{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable(model.AddressORM{}, model.EducationORM{}, model.MediaORM{}, model.SubscriptionsORM{}, model.SocialMediaORM{},
					model.DetailsORM{}, model.ExperienceORM{}, model.InvestmentORM{}, model.UserORM{}, model.ProfileORM{}, model.GroupORM{},
					model.TeamORM{}, model.TeamProfileORM{}, model.InvestorDetailORM{}, model.StartupDetailORM{}, model.SettingsORM{}, model.LoginActivityORM{},
					model.PaymentsORM{}, model.CardORM{}, model.PinORM{}, model.Privacy{}, model.NotificationORM{}, model.PostAndCommentsPushNotificationORM{},
					model.FollowingAndFollowersPushNotificationORM{}, model.DirectMessagesPushNotificationORM{}, model.EmailAndSmsPushNotificationORM{}).Error
			},
		},
	})

	err := migration.Migrate()
	if err != nil {
		return err
	}

	return nil
}
