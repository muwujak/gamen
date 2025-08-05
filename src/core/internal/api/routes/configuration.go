package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mujak27/gamen/src/core/internal/api/handlers"
)

func RegisterConfigurationRoutes(router *gin.RouterGroup, configurationHandler *handlers.ConfigurationHandler) {
	// TODO: need to think, singular or plural?
	configurationGroup := router.Group("/configurations")
	{
		configurationGroup.GET("", configurationHandler.ListConfigurations)
		configurationGroup.POST("", configurationHandler.CreateConfiguration)
		configurationGroup.GET("/:id", configurationHandler.GetConfigurationById)
		configurationGroup.PUT("/:id", configurationHandler.UpdateConfiguration)
		configurationGroup.DELETE("/:id", configurationHandler.DeleteConfiguration)
	}
	configurationTypeGroup := router.Group("/configuration-types")
	{
		configurationTypeGroup.GET("", configurationHandler.ListConfigurationTypes)
		configurationTypeGroup.GET("/:id", configurationHandler.GetConfigurationById)
	}
}
