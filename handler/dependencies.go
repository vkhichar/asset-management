package handler

import (
	"github.com/vkhichar/asset-management/repository"
	"github.com/vkhichar/asset-management/service"
)

type dependencies struct {
	userService  service.UserService
	assetService service.AssetService
	tokenService service.TokenService
	userService             service.UserService
	assetMaintenanceService service.AssetMaintenanceService
	assetService            service.AssetService
}

var deps dependencies

func InitDependencies() {
	userRepo := repository.NewUserRepository()
	assetRepo := repository.NewAssetRepository()
	assetService := service.NewAssetService(assetRepo)
	deps.assetService = assetService
	jwtTokenService := service.NewJwtService()
	deps.userService = service.NewUserService(userRepo, jwtTokenService)
	deps.tokenService = jwtTokenService
	plainTokenService := service.NewPlainTokenService()
	assetMaintainRepo := repository.NewAssetMaintainRepository()
	assetMaintenanceService := service.NewAssetForMaintenance(assetMaintainRepo)
	assetRepo := repository.NewAssetRepository()
	assetService := service.NewAssetService(assetRepo)
	eventSvc := service.NewEventService()
	userService := service.NewUserService(userRepo, plainTokenService, eventSvc)
	deps.userService = userService
	deps.assetService = assetService
	deps.assetMaintenanceService = assetMaintenanceService
}
