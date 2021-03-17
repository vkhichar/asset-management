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
	assetRepo := repository.NewAssetRepository()
	eventSvc := service.NewEventService()
	assetService := service.NewAssetService(assetRepo, eventSvc)
	jwtTokenService := service.NewJwtService()

	assetMaintainRepo := repository.NewAssetMaintainRepository()
	assetMaintenanceService := service.NewAssetForMaintenance(assetMaintainRepo)
	userService := service.NewUserService(userRepo, plainTokenService, eventSvc)
	assetMaintenanceRepo := repository.NewAssetMaintainRepository()

	deps.assetMaintenanceService = service.NewAssetForMaintenance(assetMaintenanceRepo, eventSvc)
	deps.userService = userService
	deps.assetService = assetService
	deps.tokenService = jwtTokenService
}
