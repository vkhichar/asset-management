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

	userService := service.NewUserService(userRepo, plainTokenService)
	deps.userService = userService

	assetMaintenanceService := service.NewAssetForMaintenance(assetMaintainRepo)
	deps.assetMaintenanceService = assetMaintenanceService
	assetRepo := repository.NewAssetRepository()
	assetService := service.NewAssetService(assetRepo)

	deps.userService = userService
	deps.assetService = assetService
}
