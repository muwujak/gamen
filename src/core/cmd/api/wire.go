package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mujak27/gamen/src/core/internal/api/handlers"
	"github.com/mujak27/gamen/src/core/internal/api/routes"
	kubernetes_configuration_type "github.com/mujak27/gamen/src/core/internal/extensions/configuration_types/kubernetes/services"
	restart_deployment_plugin "github.com/mujak27/gamen/src/core/internal/extensions/plugins/restart_deployment/services"
	"github.com/mujak27/gamen/src/core/internal/repository/databases"
	"github.com/mujak27/gamen/src/core/internal/services"
	"gorm.io/gorm"
)

func Wire(r *gin.Engine, db *gorm.DB) {

	configurationRepo := databases.NewConfigurationRepository(db)
	dashboardRepo := databases.NewDashboardRepository(db)
	catalogueRepo := databases.NewCatalogueRepository(db)

	userRepo := databases.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	configurationService := services.NewConfigurationService(configurationRepo, userService)
	dashboardService := services.NewDashboardService(dashboardRepo)
	catalogueService := services.NewCatalogueService(catalogueRepo, configurationService)

	configurationHandler := handlers.NewConfigurationHandler(configurationService)
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)
	catalogueHandler := handlers.NewCatalogueHandler(catalogueService)

	widgetRepo := databases.NewWidgetRepository(db)
	widgetService := services.NewWidgetService(catalogueService, configurationService, widgetRepo)
	widgetHandler := handlers.NewWidgetHandler(widgetService)

	// configuration types

	// kubernetes configuration type
	kubernetesConfigurationService := kubernetes_configuration_type.NewKubernetesConfigurationTypeService()
	configurationRepo.RegisterConfigurationTypeService(kubernetesConfigurationService.GetId(), kubernetesConfigurationService)

	// catalogues

	// restart deployment catalogue
	restartDeploymentUtilService := restart_deployment_plugin.NewUtilService()
	restartDeploymentService := restart_deployment_plugin.NewRestartDeploymentService(kubernetesConfigurationService, catalogueRepo, restartDeploymentUtilService)
	// register catalogue to static repository
	catalogueRepo.RegisterPluginFunction(restartDeploymentService.GetId(), restartDeploymentService)

	routes.RegisterRoutes(r, configurationHandler, dashboardHandler, catalogueHandler, userHandler, widgetHandler)
}
