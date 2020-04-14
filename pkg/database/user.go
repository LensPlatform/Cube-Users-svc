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

		// check if the user exists in the database base off of
		// email and username
		// Note: Email and Usernames must be unique across the entire database
		if err = tx.Where("email = ?", user.Email).Or("username = ?", user.UserName).Find(&foundUser).Error; err != nil {
			return err
		}

		// if the user already exists raise and error
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

func (db *Database) PerformTransaction(functor func(tx *gorm.DB) error) error {
	return db.Engine.Transaction(functor)
}
