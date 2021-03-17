package handler

import (
	"github.com/vkhichar/asset-management/repository"
	"github.com/vkhichar/asset-management/service"
)

type dependencies struct {
	userService  service.UserService
	assetService service.AssetService
}

var deps dependencies

func InitDependencies() {
	userRepo := repository.NewUserRepository()
	plainTokenService := service.NewPlainTokenService()
	eventService := service.NewEventService()
	assetRepo := repository.NewAssetRepository()
	assetService := service.NewAssetService(assetRepo)
	userService := service.NewUserService(userRepo, plainTokenService, eventService)
	deps.userService = userService
	deps.assetService = assetService
}
