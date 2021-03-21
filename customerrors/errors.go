package customerrors

import "errors"

var (
	UserNotExist = errors.New("User does not exist")

	ExtraError                = errors.New("invalid email")
	ErrInvalidAssetStatus     = errors.New("Invalid asset status")
	ErrInvalidEmailPassword   = errors.New("invalid email or password")
	NoUsersExist              = errors.New("No users exist at present")
	NoAssetsExist             = errors.New("No assets exist")
	MaintenanceIdDoesNotExist = errors.New("Maintenance id does no exist")
	ResponseTimeLimitExceeded = errors.New("Response Time Limit Exceeded")
	UserDoesNotExist          = errors.New("The user for this id does not exist")
	ErrMissingToken           = errors.New("missing token")
	ErrInvalidToken           = errors.New("invalid or expired token")
	ErrForbidden              = errors.New("forbidden")
	ErrBadRequest             = errors.New("bad request")
	ErrNotFound               = errors.New("not found")
	AssetDoesNotExist         = errors.New("asset for this id does not exist")
	AssetCannotBeAllocated    = errors.New("this asset is either retired or is under maintenance")
	AssetAlreadyAllocated     = errors.New("this asset is already allocated to another user")
)
