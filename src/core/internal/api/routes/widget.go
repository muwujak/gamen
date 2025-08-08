package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mujak27/gamen/src/core/internal/api/middleware"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/handler"
)

func RegisterWidgetRoutes(router *gin.RouterGroup, widgetHandler interfaces.WidgetHandler) {

	widgetGroup := router.Group("/widgets")
	widgetGroup.Use(middleware.AuthMiddleware())
	{
		widgetGroup.POST("/action", widgetHandler.Action)
	}
}
