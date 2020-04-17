package customerror

import (
	"errors"
)

var (
	UserprofileAlreadyExists  = errors.New("User profile already exists")
	UserAccountDoesNotExist   = errors.New("User account does not exist. Please create a user account first")
	SubscriptionAlreadyExists = errors.New("Subscription already exists")
)
