package handler

import (
	"github.com/vkhichar/asset-management/repository"
	"github.com/vkhichar/asset-management/service"
)

type dependencies struct {
	userService             service.UserService
	assetMaintenanceService service.AssetMaintenanceService
	assetService            service.AssetService
	assetMaintenanceService service.AssetMaintenanceService
	tokenService            service.TokenService
}

var deps dependencies

func InitDependencies() {
	userRepo := repository.NewUserRepository()
	plainTokenService := service.NewPlainTokenService()
	assetMaintainRepo := repository.NewAssetMaintainRepository()
	assetRepo := repository.NewAssetRepository()
	assetService := service.NewAssetService(assetRepo)
	eventSvc := service.NewEventService()
	userService := service.NewUserService(userRepo, plainTokenService, eventSvc)

	deps.userService = userService
	deps.assetService = assetService
	deps.assetMaintenanceService = assetMaintenanceService
	deps.tokenService = plainTokenService
	assetMaintenanceRepo := repository.NewAssetMaintainRepository()
	deps.assetMaintenanceService = service.NewAssetForMaintenance(assetMaintenanceRepo)
}
