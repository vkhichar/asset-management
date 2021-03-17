package customerrors

import "errors"

var (
	ErrInvalidEmailPassword = errors.New("invalid email or password")
	NoUsersExist            = errors.New("No users exist at present")
	AssetAlreadyDeleted     = errors.New("Asset Already Deleted")
	UserDoesNotExist        = errors.New("The user for this id does not exist")
	NoAssetsExist           = errors.New("No assets exist")
	NoUserExistForDelete    = errors.New("No user present by this Id")
)
