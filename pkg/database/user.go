package database

import (
	"context"
	"errors"

	model "github.com/LensPlatform/micro/pkg/models/proto"
	"github.com/jinzhu/gorm"
)

func (db *Database) CreateUser(user model.User) error {
	functor := func(tx *gorm.DB) error {
		var foundUser model.UserORM
		ctx := context.TODO()
		err := user.Validate()
		if err != nil {
			return err
		}

		// check if the user exists in the database based off of
		// email and username
		// Note: Email and Usernames must be unique across the entire database
		if err = tx.Where("email = ?", user.Email).Or("username = ?", user.UserName).Find(&foundUser).Error; err != nil {
			return err
		}

		// if the user already exists raise an error
		if foundUser.UserName != "" || foundUser.Email != "" {
			return errors.New("user already exists")
		}

		// convert the user field to orm
		userOrm, err := user.ToORM(ctx)
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
	var user model.UserORM

	if err := db.Engine.Where("id = ?", user_id).Find(&user).Error; err != nil {
		return err, &model.User{}
	}

	// convert the obtained user ORM object to a user object and validate all fields are there
	ctx := context.TODO()
	userObj, err := user.ToPB(ctx)
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
		err, exists := db.DoesUserExists(user_id, "", "")
		if !exists && err != nil {
			db.Logger.Log(err.Error())
			return err
		}

		// query the database for user by user id
		if err = tx.Where("id = ?", user_id).Find(&userOrm).Error; err != nil {
			return err
		}

		ctx := context.TODO()
		userObj, err := userOrm.ToPB(ctx)
		if err != nil {
			return err
		}

		// validate that the user has all the fields of interest
		if err = userObj.Validate(); err != nil {
			return err
		}

		// validate that the profile has al fields of interest as well
		if err = profile.Validate(); err != nil {
			return err
		}

		// check that the same profile does not already exist
		err, profileExsts := db.DoesUserProfileExists(user_id)
		if err != nil && !profileExsts {
			return nil
		}

		// update the user ORM object with the profile ORM object
		profileOrm, err := profile.ToORM(ctx)
		if err != nil {
			return nil
		}

		// Updates all fields in a user entity in the database
		if err = tx.Model(userOrm).Updates(model.UserORM{ProfileId: &profileOrm}).Error; err != nil {
			return err
		}

		return nil
	}

	return db.PerformTransaction(functor)
}

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
		if user.Id == 0 {
			return errors.New("User Does Not Exist")
		}

		return nil
	}

	err := db.PerformTransaction(functor)
	if err != nil {
		db.Logger.Log(err.Error())
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
			return errors.New("Profile Of Interest Does Not Exist")
		}

		return nil
	}

	err := db.PerformTransaction(functor)
	if err != nil {
		return err, false
	}

	return nil, true
}

func (db *Database) PerformTransaction(functor func(tx *gorm.DB) error) error {
	return db.Engine.Transaction(functor)
}
