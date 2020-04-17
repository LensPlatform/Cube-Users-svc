package database

import (
	"context"
	"errors"

	"github.com/LensPlatform/micro/user-microservice/pkg/helper/types"
	model "github.com/LensPlatform/micro/user-microservice/pkg/models/proto"
	"github.com/jinzhu/gorm"
)

func (db *Database) CreateUser(user model.User) error {
	functor := func(tx *gorm.DB) error {
		var userOrm model.UserORM

		// check if the user exists in the database based off of
		// email and username
		// Note: Email and Usernames must be unique across the entire database
		if err := tx.Where("email = ? AND username = ?", user.Email, user.UserName).Find(&userOrm).Error; err != nil {
			return err
		}

		// if the user already exists raise an error
		if userOrm.UserName != "" || userOrm.Email != "" || userOrm.Id == 0 {
			return errors.New("user already exists")
		}

		// convert the user field to orm
		userOrm, err := user.ToORM(context.TODO())
		if err != nil {
			return err
		}

		// save the user to the database
		if err := tx.Create(userOrm).Error; err != nil {
			return err
		}

		return nil
	}

	return db.PerformTransaction(functor)
}

func (db *Database) GetUserById(user_id int32) (error, *model.User) {
	// query the database for the user of interest
	var userOrm model.UserORM

	if err := db.Engine.Where("id = ?", user_id).Find(&userOrm).Error; err != nil {
		return err, &model.User{}
	}

	// convert the obtained user ORM object to a user object and validate all fields are there
	userObj, err := userOrm.ToPB(context.TODO())
	if err != nil {
		return err, &model.User{}
	}

	// perform field validation
	if err = userObj.Validate(); err != nil {
		return err, &model.User{}
	}

	return nil, &userObj
}

func (db *Database) CreateUserProfile(user_id int32, profile model.Profile) error {
	functor := func(tx *gorm.DB) error {
		var userOrm model.UserORM

		// check that the user exists
		err, exists := db.DoesUserExists(user_id, "", "")
		if !exists {
			return err
		}

		// query the database for user by user id
		if err = tx.Where("id = ?", user_id).Find(&userOrm).Error; err != nil {
			return err
		}

		if userOrm.Id == 0 || userOrm.UserName == "" || userOrm.LastName == "" {
			return errors.New("user does not exist")
		}

		// check that the same profile does not already exist
		err, profileExists := db.DoesUserProfileExists(user_id)
		if profileExists {
			return err
		}

		// update the user ORM object with the profile ORM object
		profileOrm, err := profile.ToORM(context.TODO())
		if err != nil {
			return err
		}

		// Updates only the relevant fields of interest in a user entity in the database
		if err = tx.Model(userOrm).Updates(model.UserORM{ProfileId: &profileOrm}).Error; err != nil {
			return err
		}

		return nil
	}

	return db.PerformTransaction(functor)
}

func (db *Database) CreateUserSubscription(user_id int32, subscription model.Subscriptions) error {
	// There are no assumptions about the state of the database with respect to the current
	// subscription. Hence validation should be performed
	// prior to the subscription being created

	functor := func(tx *gorm.DB) error {
		var (
			userOrm       model.UserORM
			subscriptions []*model.SubscriptionsORM
		)

		// validate the subscription object
		if subscription.SubscriptionName == "" ||
			subscription.StartDate == nil ||
			subscription.EndDate == nil ||
			subscription.SubscriptionStatus == "" {
			return errors.New("invalid subscription")
		}

		// check and make sure the  user account with the specified userid exists
		err, exist := db.DoesUserExists(user_id, types.Empty, types.Empty)
		if !exist {
			return err
		}

		// obtain the user based off of user_id
		if err := tx.Where("id = ?", user_id).Find(&userOrm).Error; err != nil {
			return err
		}

		if userOrm.Id == 0 {
			return errors.New("user account does not exist")
		}

		// convert the subscription object to an ORM type
		subscriptionOrm, err := subscription.ToORM(context.TODO())
		if err != nil {
			return err
		}

		err, exist = db.DoesSubscriptionExist(subscriptionOrm.SubscriptionName, user_id)
		if exist {
			// iterate over all of the user's subscriptions and update the status of the subscription to active
			for _, subsc := range userOrm.SubscriptionsId {
				if subsc.SubscriptionName == subscriptionOrm.SubscriptionName {
					// activate the subscription if it is not already active
					subsc.IsActive = true
					subsc.EndDate = subscriptionOrm.EndDate
				}
				subscriptions = append(subscriptions, subsc)
			}
			userOrm.SubscriptionsId = subscriptions
		} else {
			// just add the new subscription to the user orm
			subscriptions = append(subscriptions, &subscriptionOrm)
			userOrm.SubscriptionsId = subscriptions
		}

		// save the user
		if err = tx.Save(userOrm).Error; err != nil {
			return err
		}

		return nil
	}

	return db.PerformTransaction(functor)
}

func (db *Database) UpdateUser(user_id int32, user model.User) (error, *model.User) {

	transaction := func(tx *gorm.DB) (error, interface{}) {
		// first and foremost we check for the existence of the user
		err, exist := db.DoesUserExists(user_id, types.Empty, types.Empty)
		if !exist {
			return err, &model.User{}
		}

		// convert the user to an ORM type
		userOrm, err := user.ToORM(context.TODO())
		if err != nil {
			return err, &model.User{}
		}

		// update the actual user in the database
		if err := tx.Save(&userOrm).Error; err != nil {
			tx.Rollback()
			return err, &model.User{}
		}

		return nil, user
	}

	err, output := db.PerformComplexTransaction(transaction)

	if err != nil {
		return err, &model.User{}
	}

	updatedUser := output.(model.User)
	return nil, &updatedUser
}

func (db *Database) UpdateUserSubscription(user_id int32, subscription_id int32, subscription model.Subscriptions) (error, *model.Subscriptions) {
	transaction := func(tx *gorm.DB) (error, interface{}) {
		var (
			userOrm              model.UserORM
			updatedSubscriptions []*model.SubscriptionsORM
		)
		// first and foremost we check for the existence of the user account
		// and the subscription
		err, exist := db.DoesSubscriptionExist(subscription.SubscriptionName, user_id)
		if !exist {
			return err, &model.Subscriptions{}
		}

		// obtain the user from the database
		if err := tx.Where("id = ?", user_id).Find(userOrm).Error; err != nil {
			return err, &model.Subscriptions{}
		}

		// convert the profile to an ORM type
		subscriptionOrm, err := subscription.ToORM(context.TODO())
		if err != nil {
			return err, &model.Subscriptions{}
		}

		for _, oldSubscription := range userOrm.SubscriptionsId {
			if oldSubscription.SubscriptionName == subscriptionOrm.SubscriptionName {
				updatedSubscriptions = append(updatedSubscriptions, &subscriptionOrm)
			} else {
				updatedSubscriptions = append(updatedSubscriptions, oldSubscription)
			}
		}

		// update the subscription list tied to the user
		userOrm.SubscriptionsId = updatedSubscriptions

		// update the actual user in the database
		if err := tx.Save(&userOrm).Error; err != nil {
			tx.Rollback()
			return err, &model.Profile{}
		}

		return nil, &subscription
	}

	err, output := db.PerformComplexTransaction(transaction)

	if err != nil {
		return err, &model.Subscriptions{}
	}

	updatedSubscription := output.(model.Subscriptions)
	return nil, &updatedSubscription
}

func (db *Database) UpdateUserProfile(user_id, profile_id int32, profile model.Profile) (error, *model.Profile) {
	transaction := func(tx *gorm.DB) (error, interface{}) {
		var userOrm model.UserORM

		// first and foremost we check for the existence of the user account
		// and the profile
		err, exist := db.DoesUserProfileExists(user_id)
		if !exist {
			return err, &model.Profile{}
		}

		// obtain the user from the database
		if err := tx.Where("id = ?", user_id).Find(userOrm).Error; err != nil {
			return err, &model.Profile{}
		}

		// convert the profile to an ORM type
		profileOrm, err := profile.ToORM(context.TODO())
		if err != nil {
			return err, &model.Profile{}
		}

		// update the profile tied to the user
		userOrm.ProfileId = &profileOrm

		// update the actual user in the database
		if err := tx.Save(&userOrm).Error; err != nil {
			tx.Rollback()
			return err, &model.Profile{}
		}

		return nil, profile
	}

	err, output := db.PerformComplexTransaction(transaction)

	if err != nil {
		return err, &model.Profile{}
	}

	updatedProfile := output.(model.Profile)
	return nil, &updatedProfile
}

func (db *Database) DeleteUser(user_id int32) error {
	transaction := func(tx *gorm.DB) error {
		var userOrm model.UserORM

		err, exist := db.DoesUserExists(user_id, types.Empty, types.Empty)
		if !exist {
			return err
		}

		// obtain user of interest
		if err := tx.First(&userOrm, user_id).Error; err != nil {
			return err
		}

		// delete user from dab
		if err = tx.Where("id =?", user_id).Delete(&userOrm).Error; err != nil {
			return err
		}

		return nil
	}

	return db.PerformTransaction(transaction)
}

func (db *Database) DeleteUserProfile(user_id, profile_id int32) error {
	transaction := func(tx *gorm.DB) error {
		var userOrm model.UserORM

		// check the user of interest has a profile to even delete
		err, exist := db.DoesUserProfileExists(user_id)
		if !exist {
			return err
		}

		// obtain user of interest
		if err := tx.First(&userOrm, user_id).Error; err != nil {
			return err
		}

		// obtain the profile from the user account returned from the database
		profileOrm := *userOrm.ProfileId

		if profileOrm.Id != profile_id {
			return errors.New("profiled id doest not exist for this user.")
		}

		// delete user from dab
		if err = tx.Where("id = ?", profile_id).Delete(&profileOrm).Error; err != nil {
			return err
		}

		return nil
	}

	return db.PerformTransaction(transaction)
}

func (db *Database) DeleteUserSubscription(user_id, subscription_id int32) error {
	transaction := func(tx *gorm.DB) error {
		// check the user of interest exists
		err, exist := db.DoesUserExists(user_id, types.Empty, types.Empty)
		if !exist {
			return err
		}

		err, exist = db.DoesSubscriptionExistById(user_id, subscription_id)
		if !exist {
			return err
		}

		// delete from subscriptions table in db where id == subscriptions id
		if err = tx.Where("id = ?", subscription_id).Delete(&model.SubscriptionsORM{}).Error; err != nil {
			return err
		}

		return nil
	}

	return db.PerformTransaction(transaction)
}

func (db *Database) PerformTransaction(functor func(tx *gorm.DB) error) error {
	return db.Engine.Transaction(functor)
}

func (db *Database) PerformComplexTransaction(transaction func(tx *gorm.DB) (error, interface{})) (error, interface{}) {
	tx := db.Engine.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err, nil
	}

	err, result := transaction(tx)
	if err != nil {
		return err, nil
	}

	return tx.Commit().Error, result
}
