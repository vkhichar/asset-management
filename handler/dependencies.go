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
	eventSvc := service.NewEventService()
	assetService := service.NewAssetService(assetRepo, eventSvc)
	jwtTokenService := service.NewJwtService()
	plainTokenService := service.NewPlainTokenService()


	eventSvc := service.NewEventService()

	assetMaintainRepo := repository.NewAssetMaintainRepository()
	assetMaintenanceService := service.NewAssetForMaintenance(assetMaintainRepo)

	userService := service.NewUserService(userRepo, plainTokenService, eventSvc)
	deps.userService = userService
	deps.assetService = assetService
	deps.tokenService = jwtTokenService
	deps.tokenService = plainTokenService
	assetMaintenanceRepo := repository.NewAssetMaintainRepository()
	eventService := service.NewEventService()
	deps.assetMaintenanceService = service.NewAssetForMaintenance(assetMaintenanceRepo, eventService)
}
