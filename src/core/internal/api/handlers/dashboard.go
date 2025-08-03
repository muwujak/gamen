package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mujak27/gamen/src/core/internal/services"
)

type DashboardHandler struct {
	service *services.DashboardService
}

func NewDashboardHandler(service *services.DashboardService) *DashboardHandler {
	return &DashboardHandler{service: service}
}

func (dashboardHandler *DashboardHandler) GetCurrentUserDashboard(ctx *gin.Context) {
	dashboard, err := dashboardHandler.service.GetCurrentUserDashboard()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve dashboard"})
		return
	}
	ctx.JSON(200, gin.H{"dashboard": dashboard})
}
