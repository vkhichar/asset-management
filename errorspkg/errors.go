package errorspkg

import "errors"

var ErrInvalidEmailPassword = errors.New("invalid email or password")

var NoUsersExist = errors.New("No users exist at present")
