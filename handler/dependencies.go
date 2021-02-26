package handler

import (
	"github.com/vkhichar/asset-management/service"
)

type dependencies struct {
	userService service.UserService
}

var deps dependencies

func InitDependencies() {
	userService := service.NewUserService()
	deps.userService = userService
}
