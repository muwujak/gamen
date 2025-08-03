package routes

import (
	"github.com/gin-gonic/gin"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/handler"
)

// TODO: think, delete?
func RegisterPluginRoutes(router *gin.RouterGroup, pluginHandler interfaces.PluginHandler) {

	id := pluginHandler.GetId()

	catalogueGroup := router.Group("/catalogue/" + id.String())
	{
		catalogueGroup.GET("fetch", pluginHandler.Fetch)
		catalogueGroup.POST("action", pluginHandler.Action)
		catalogueGroup.GET("get", pluginHandler.Get)
	}
}
