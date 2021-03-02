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

	assetRepo := repository.NewAssetRepository()
	assetService := service.NewAssetService(assetRepo)
	deps.assetService = assetService
	jwtTokenService := service.NewJwtService()
	deps.userService = service.NewUserService(userRepo, jwtTokenService)
}
