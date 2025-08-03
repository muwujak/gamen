package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mujak27/gamen/src/core/internal/api/handlers"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/handler"
)

func RegisterRoutes(
	router *gin.Engine,
	configurationHandler *handlers.ConfigurationHandler,
	dashboardHandler *handlers.DashboardHandler,
	catalogueHandler *handlers.CatalogueHandler,
) {
	apiV1 := router.Group("/api/v1")
	RegisterConfigurationRoutes(apiV1, configurationHandler)
	RegisterDashboardRoutes(apiV1, dashboardHandler)
	RegisterCatalogueRoutes(apiV1, catalogueHandler)
}

func RegisterPluginListRoute(
	router *gin.Engine,
	pluginHandler []interfaces.PluginHandler,
) {
	apiV1 := router.Group("/api/v1")
	for _, pluginHandler := range pluginHandler {
		RegisterPluginRoutes(apiV1, pluginHandler)
	}
}
