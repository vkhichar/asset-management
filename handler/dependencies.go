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
	assetMaintenanceService := service.NewAssetForMaintenance(assetMaintainRepo, eventSvc)
	deps.assetMaintenanceService = assetMaintenanceService
	assetRepo := repository.NewAssetRepository()
	assetService := service.NewAssetService(assetRepo, eventSvc)
	jwtTokenService := service.NewJwtService()
	userService := service.NewUserService(userRepo, plainTokenService, eventSvc)
	deps.userService = userService
	deps.assetService = assetService
	deps.tokenService = jwtTokenService
}
