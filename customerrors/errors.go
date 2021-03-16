package customerrors

import "errors"

var (
	ErrInvalidEmailPassword = errors.New("invalid email or password")

	NoUsersExist = errors.New("No users exist at present")

	NoAssetsExist       = errors.New("No asset found")
	AssetAlreadyDeleted = errors.New("Asset Already Deleted")
)
