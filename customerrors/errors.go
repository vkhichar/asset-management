package customerrors

import "errors"

var (
	ErrInvalidEmailPassword = errors.New("invalid email or password")
	NoUsersExist = errors.New("No users exist at present")
	NoAssetsExist = errors.New("No assets exist")
	NoMaintenanceActivitesExist = errors.New("No maintenance activites exist")
	MaintenanceIdDoesNotExist = errors.New("Maintenance id does no exist")
	AssetAlreadyDeleted     = errors.New("Asset Already Deleted")
	UserDoesNotExist        = errors.New("The user for this id does not exist")
	NoUserExistForDelete    = errors.New("No user present by this Id")
)
