package customerrors

import "errors"

var (
	ErrInvalidEmailPassword   = errors.New("invalid email or password")
	NoUsersExist              = errors.New("No users exist at present")
	NoAssetsExist             = errors.New("No assets exist")
	MaintenanceIdDoesNotExist = errors.New("Maintenance id does no exist")
	UserDoesNotExist          = errors.New("The user for this id does not exist")
	NoUserExistForDelete      = errors.New("No user present by this Id")
	ErrMissingToken           = errors.New("missing token")
	ErrInvalidToken           = errors.New("invalid or expired token")
	ErrForbidden              = errors.New("forbidden")
	ErrBadRequest             = errors.New("bad request")
	ErrNotFound               = errors.New("not found")
)
