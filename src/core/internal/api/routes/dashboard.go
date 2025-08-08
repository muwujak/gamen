package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mujak27/gamen/src/core/internal/api/handlers"
	"github.com/mujak27/gamen/src/core/internal/api/middleware"
)

func RegisterDashboardRoutes(router *gin.RouterGroup, dashboardHandler *handlers.DashboardHandler) {
	dashboardGroup := router.Group("/dashboard")
	dashboardGroup.Use(middleware.AuthMiddleware())
	{
		dashboardGroup.GET("", dashboardHandler.GetCurrentUserDashboard)
	}
}
