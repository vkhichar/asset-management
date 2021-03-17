package customerrors

import "errors"

var (
	ErrInvalidEmailPassword   = errors.New("invalid email or password")
	NoUsersExist              = errors.New("No users exist at present")
	NoAssetsExist             = errors.New("No assets exist")
	MaintenanceIdDoesNotExist = errors.New("Maintenance id does no exist")
	UserDoesNotExist          = errors.New("The user for this id does not exist")
	NoUserExistForDelete      = errors.New("No user present by this Id")
)
