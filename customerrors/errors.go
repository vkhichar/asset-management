package customerrors

import "errors"

var (
	ErrInvalidEmailPassword = errors.New("invalid email or password")

	NoUsersExist = errors.New("No users exist at present")

	NoAssetsExist = errors.New("No assets exist")

	UserNotExist = errors.New("User does not exist")

	ExtraError = errors.New("invalid email")
)
