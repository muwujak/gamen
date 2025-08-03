package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mujak27/gamen/src/core/internal/api/dto"
	"github.com/mujak27/gamen/src/core/internal/services"
)

type WidgetHandler struct {
	service *services.WidgetService
}

func NewWidgetHandler(service *services.WidgetService) *WidgetHandler {
	return &WidgetHandler{
		service: service,
	}
}

func (h *WidgetHandler) Action(c *gin.Context) {
	payload := dto.WidgetActionPayload{}
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.service.Action(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Widget action performed successfully"})
}
