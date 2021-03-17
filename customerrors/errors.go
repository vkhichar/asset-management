package customerrors

import "errors"

var (
	ErrInvalidEmailPassword = errors.New("invalid email or password")

	NoUsersExist = errors.New("No users exist at present")

	NoAssetsExist = errors.New("No assets exist")

	NoMaintenanceActivitesExist = errors.New("No maintenance activites exist")

	MaintenanceIdDoesNotExist = errors.New("Maintenance id does no exist")

	ResponseTimeLimitExceeded = errors.New("Response Time Limit Exceeded")
)
