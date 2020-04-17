package database

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"

	"github.com/LensPlatform/micro/user-microservice/pkg/helper/types"
	model "github.com/LensPlatform/micro/user-microservice/pkg/models/proto"
)

func (db *Database) DoesUserExists(user_id int32, username, email string) (error, bool) {
	functor := func(tx *gorm.DB) error {
		// check if the user exists based on username, email and id
		var user model.UserORM

		if user_id != 0 && username != "" && email != "" {
			if err := tx.Where("id = ? AND email = ? AND username = ?", user_id, email, username).Find(&user).Error; err != nil {
				return err
			}
		} else if user_id != 0 && username == "" && email == "" {
			// user name and email is empty so obtain user by querying for user id
			if err := tx.Where("id = ?", user_id).Find(&user).Error; err != nil {
				return err
			}
		} else {
			return errors.New("Invalid input parameters")
		}

		// Convert the userORM to user object and perform validation checks
		ctx := context.TODO()
		userObj, err := user.ToPB(ctx)
		if err != nil {
			return err
		}

		// Actually perfom field validation
		if err := userObj.Validate(); err != nil {
			return err
		}

		// Validate that the obtained user is in fact populated
		if user.Id == 0 || user.Email == "" || user.UserName == "" {
			return errors.New("user does not exist")
		}

		return nil
	}

	err := db.PerformTransaction(functor)
	if err != nil {
		_ = db.Logger.Log(err.Error())
		return err, false
	}

	return nil, true
}

func (db *Database) DoesSubscriptionExist(subscription_name string, user_id int32) (error, bool) {
	functor := func(tx *gorm.DB) error {
		var userOrm model.UserORM
		var subscriptions []*model.SubscriptionsORM

		err, exists := db.DoesUserExists(user_id, types.Empty, types.Empty)
		if !exists {
			return err
		}

		// obtain the user to which this subscription is supposed to belong to
		if err := tx.Where("id = ?", user_id).Find(&userOrm).Error; err != nil {
			return err
		}

		// obtain all the subscriptions from the user
		subscriptions = userOrm.SubscriptionsId

		// iterate over all such subscriptions searching for the subscription name of interest
		for _, subscription := range subscriptions {
			if subscription.SubscriptionName == subscription_name {
				return nil
			}
		}

		return errors.New("subscription does not exist")
	}

	if err := db.PerformTransaction(functor); err != nil {
		return err, false
	}

	return nil, true
}

func (db *Database) DoesSubscriptionExistById(user_id, subscription_id int32) (error, bool) {
	functor := func(tx *gorm.DB) error {
		var userOrm model.UserORM
		var subscriptions []*model.SubscriptionsORM

		err, exists := db.DoesUserExists(user_id, types.Empty, types.Empty)
		if !exists {
			return err
		}

		// obtain the user to which this subscription is supposed to belong to
		if err := tx.Where("id = ?", user_id).Find(&userOrm).Error; err != nil {
			return err
		}

		// obtain all the subscriptions from the user
		subscriptions = userOrm.SubscriptionsId

		// iterate over all such subscriptions searching for the subscription id of interest
		for _, subscription := range subscriptions {
			if subscription.Id == subscription_id {
				return nil
			}
		}

		return errors.New("subscription does not exist")
	}

	if err := db.PerformTransaction(functor); err != nil {
		return err, false
	}

	return nil, true
}

func (db *Database) DoesUserProfileExists(user_id int32) (error, bool) {
	functor := func(tx *gorm.DB) error {
		var profile model.ProfileORM
		var user model.UserORM

		// preliminarily, check if the user account exists
		e0, exists := db.DoesUserExists(user_id, "", "")
		if !exists {
			db.Logger.Log(e0.Error())
			return errors.New("User account does not exist. Create an account before a profile")
		}

		// if the user account does exist, query the user table for the user of interest
		// and return an error if one arises
		if e0 = tx.Where("id = ?", user_id).Find(&user).Error; e0 != nil {
			return e0
		}

		ctx := context.TODO()

		// convert the returned userORM to a user object and validate that all
		// fields of interest are present
		userProfile, e0 := profile.ToPB(ctx)
		if e0 != nil {
			return e0
		}

		// if the user object witholds invalid fields, return an error
		if e0 = userProfile.Validate(); e0 != nil {
			return e0
		}

		// after validating the user object's fields, obtain the profile of interest from the user table
		profile = *user.ProfileId

		// Return an error if the obtained profile object witholds a nil id
		if profile.Id == 0 {
			return errors.New("profile of interest does not exist")
		}

		return nil
	}

	err := db.PerformTransaction(functor)
	if err != nil {
		return err, false
	}

	return nil, true
}
