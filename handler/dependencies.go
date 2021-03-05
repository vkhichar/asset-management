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
	assetRepo := repository.NewAssetRepository()
	assetService := service.NewAssetService(assetRepo)
	userService := service.NewUserService(userRepo, plainTokenService)
	deps.userService = userService
	deps.assetService = assetService
}
