package helper

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	model "github.com/LensPlatform/micro/user-microservice/pkg/models/proto"
)

func ValidateAndHashPassword(currentuser model.User) (user model.User, err error) {
	// check if confirmed password and actual password match
	if currentuser.Password != currentuser.PasswordConfirmed {
		return currentuser, errors.New("password and confirmed password are different")
	}

	//  hash password
	hashedPassword, err := hashAndSalt([]byte(currentuser.Password))
	if err != nil {
		return currentuser, err
	}
	// reset the hashed passwords
	currentuser.Password = hashedPassword
	currentuser.PasswordConfirmed = hashedPassword

	return currentuser, nil
}

func hashAndSalt(pwd []byte) (string, error) {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash), nil
}
