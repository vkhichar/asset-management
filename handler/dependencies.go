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
	assetAllocationService  service.AssetAllocationService
}

var deps dependencies

func InitDependencies() {
	userRepo := repository.NewUserRepository()
	plainTokenService := service.NewPlainTokenService()
	assetMaintainRepo := repository.NewAssetMaintainRepository()
	eventSvc := service.NewEventService()
	assetRepo := repository.NewAssetRepository()
	assetService := service.NewAssetService(assetRepo, eventSvc)
	assetMaintenanceService := service.NewAssetForMaintenance(assetMaintainRepo, eventSvc, assetRepo)
	deps.assetMaintenanceService = assetMaintenanceService
	userService := service.NewUserService(userRepo, plainTokenService, eventSvc)
	deps.userService = userService
	deps.assetService = assetService
	deps.tokenService = plainTokenService
	assetAllocationsRepo := repository.NewAssetAllocationRepository(assetRepo)
	assetAllocationService := service.NewAssetAllocationService(assetAllocationsRepo)
	deps.assetAllocationService = assetAllocationService
}
