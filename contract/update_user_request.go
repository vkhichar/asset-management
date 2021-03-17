package contract

type UpdateUserRequest struct {
	Name     *string `json:"name"`
	Password *string `json:"password"`
}
