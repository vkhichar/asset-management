package contract

import "github.com/vkhichar/asset-management/domain"

type ListUserResponse struct {
	Users []domain.UserList `json:"users[]"`
}
