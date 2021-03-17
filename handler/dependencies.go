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
	assetRepo := repository.NewAssetRepository()
	assetService := service.NewAssetService(assetRepo)

	jwtTokenService := service.NewJwtService()
	plainTokenService := service.NewPlainTokenService()

	eventSvc := service.NewEventService()
	userService := service.NewUserService(userRepo, plainTokenService, eventSvc)

	deps.userService = userService
	deps.assetService = assetService
	deps.tokenService = jwtTokenService
	deps.tokenService = plainTokenService
	assetMaintenanceRepo := repository.NewAssetMaintainRepository()
	eventService := service.NewEventService()
	deps.assetMaintenanceService = service.NewAssetForMaintenance(assetMaintenanceRepo, eventService)
}
