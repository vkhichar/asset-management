package handler

import (
	"github.com/vkhichar/asset-management/repository"
	"github.com/vkhichar/asset-management/service"
)

type dependencies struct {
	userService             service.UserService
	assetMaintenanceService service.AssetMaintenanceService
	assetService            service.AssetService
}

var deps dependencies

func InitDependencies() {
	userRepo := repository.NewUserRepository()
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
