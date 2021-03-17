package handler

import (
	"github.com/vkhichar/asset-management/repository"
	"github.com/vkhichar/asset-management/service"
)

type dependencies struct {
	userService             service.UserService
	assetService            service.AssetService
<<<<<<< HEAD
	assetMaintenanceService service.AssetMaintenanceService
	tokenService            service.TokenService
=======
	tokenService            service.TokenService
	assetMaintenanceService service.AssetMaintenanceService
>>>>>>> 76375e324ca621f1db6d1abfc5d5b306165d91f1
}

var deps dependencies

func InitDependencies() {
	userRepo := repository.NewUserRepository()
	assetRepo := repository.NewAssetRepository()
	assetService := service.NewAssetService(assetRepo)

	jwtTokenService := service.NewJwtService()
	plainTokenService := service.NewPlainTokenService()

	assetMaintainRepo := repository.NewAssetMaintainRepository()
	assetRepo := repository.NewAssetRepository()
	assetService := service.NewAssetService(assetRepo)
	eventSvc := service.NewEventService()
	userService := service.NewUserService(userRepo, plainTokenService, eventSvc)

	deps.userService = userService
	deps.assetService = assetService
	deps.tokenService = jwtTokenService
	deps.assetMaintenanceService = assetMaintenanceService
	deps.tokenService = plainTokenService
	assetMaintenanceRepo := repository.NewAssetMaintainRepository()
	eventService := service.NewEventService()
	deps.assetMaintenanceService = service.NewAssetForMaintenance(assetMaintenanceRepo, eventService)
}
