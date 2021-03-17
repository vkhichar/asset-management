package handler

import (
	"github.com/vkhichar/asset-management/repository"
	"github.com/vkhichar/asset-management/service"
)

type dependencies struct {
	userService             service.UserService
	assetService            service.AssetService
	tokenService            service.TokenService
	assetMaintenanceService service.AssetMaintenanceService
}

var deps dependencies

func InitDependencies() {
	userRepo := repository.NewUserRepository()

	plainTokenService := service.NewPlainTokenService()
	eventService := service.NewEventService()
	assetRepo := repository.NewAssetRepository()
	assetService := service.NewAssetService(assetRepo)
	userService := service.NewUserService(userRepo, plainTokenService, eventService)
	assetRepo := repository.NewAssetRepository()
	assetService := service.NewAssetService(assetRepo)

	jwtTokenService := service.NewJwtService()
	plainTokenService := service.NewPlainTokenService()

	assetMaintainRepo := repository.NewAssetMaintainRepository()
	assetMaintenanceService := service.NewAssetForMaintenance(assetMaintainRepo)
	eventSvc := service.NewEventService()
	userService := service.NewUserService(userRepo, plainTokenService, eventSvc)

	deps.userService = userService
	deps.assetService = assetService
	deps.tokenService = jwtTokenService
	deps.assetMaintenanceService = assetMaintenanceService
}
