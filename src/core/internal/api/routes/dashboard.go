package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mujak27/gamen/src/core/internal/api/handlers"
)

func RegisterDashboardRoutes(router *gin.RouterGroup, dashboardHandler *handlers.DashboardHandler) {
	dashboardGroup := router.Group("/dashboard")
	{
		dashboardGroup.GET("", dashboardHandler.GetCurrentUserDashboard)
	}
}
