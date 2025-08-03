package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mujak27/gamen/src/core/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/mujak27/gamen/src/core/internal/api/handlers"
	"github.com/mujak27/gamen/src/core/internal/api/routes"
	kubernetes_configuration_type "github.com/mujak27/gamen/src/core/internal/extensions/configuration_types/kubernetes/services"
	restart_deployment_plugin "github.com/mujak27/gamen/src/core/internal/extensions/plugins/restart_deployment/services"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/handler"
	"github.com/mujak27/gamen/src/core/internal/models"
	"github.com/mujak27/gamen/src/core/internal/repository/databases"
	"github.com/mujak27/gamen/src/core/internal/services"

	// _ "github.com/mujak27/gamen/src/util"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		10.227.0.249:8987
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

func main() {
	// TODO: add configuration to database
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// TODO: move migration and seeding to separate file
	db.AutoMigrate(&models.Configuration{})
	db.AutoMigrate(&models.ConfigurationType{})

	configurationRepo := databases.NewConfigurationRepository(db)
	dashboardRepo := databases.NewDashboardRepository(db)
	catalogueRepo := databases.NewCatalogueRepository(db)

	configurationService := services.NewConfigurationService(configurationRepo)
	dashboardService := services.NewDashboardService(dashboardRepo)
	catalogueService := services.NewCatalogueService(catalogueRepo, configurationService)

	configurationHandler := handlers.NewConfigurationHandler(configurationService)
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)
	catalogueHandler := handlers.NewCatalogueHandler(catalogueService)

	// configuration type
	kubernetesConfigurationService := kubernetes_configuration_type.NewKubernetesConfigurationService()

	// catalogues
	var pluginHandlerList []interfaces.PluginHandler

	// restart deployment catalogue
	restartDeploymentUtilService := restart_deployment_plugin.NewUtilService()
	restartDeploymentService := restart_deployment_plugin.NewRestartDeploymentService(kubernetesConfigurationService, catalogueRepo, restartDeploymentUtilService)
	restartDeploymentHandler := handlers.NewPluginHandler(restartDeploymentService)
	// register routing
	// TODO: remove plugin handler
	pluginHandlerList = append(pluginHandlerList, restartDeploymentHandler)
	// register catalogue to static repository
	catalogueRepo.RegisterPluginFunction(restartDeploymentService.GetId(), restartDeploymentService)

	r := gin.Default()
	routes.RegisterRoutes(r, configurationHandler, dashboardHandler, catalogueHandler)
	routes.RegisterPluginListRoute(r, pluginHandlerList)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8987")
}
