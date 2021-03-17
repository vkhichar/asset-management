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
	assetMaintainRepo := repository.NewAssetMaintainRepository()
	eventSvc := service.NewEventService()

	//userService := service.NewUserService(userRepo, plainTokenService)
	//deps.userService = userService

	assetMaintenanceService := service.NewAssetForMaintenance(assetMaintainRepo, eventSvc)
	deps.assetMaintenanceService = assetMaintenanceService
	assetRepo := repository.NewAssetRepository()
	//eventSvc := service.NewEventService()
	assetService := service.NewAssetService(assetRepo, eventSvc)
	jwtTokenService := service.NewJwtService()
	//plainTokenService := service.NewPlainTokenService()
	userService := service.NewUserService(userRepo, plainTokenService, eventSvc)
	assetMaintenanceRepo := repository.NewAssetMaintainRepository()

	deps.assetMaintenanceService = service.NewAssetForMaintenance(assetMaintenanceRepo, eventSvc)
	deps.userService = userService
	deps.assetService = assetService
	deps.tokenService = jwtTokenService
}
