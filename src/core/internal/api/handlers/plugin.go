package handlers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/api/dto"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/services"
)

type PluginHandler struct {
	PluginService interfaces.PluginService
}

func NewPluginHandler(
	restartDeploymentService interfaces.PluginService,
) *PluginHandler {
	return &PluginHandler{
		PluginService: restartDeploymentService,
	}
}

// TODO: improve routing key id system
func (h *PluginHandler) GetId() uuid.UUID {
	return h.PluginService.GetId()
}

func (h *PluginHandler) Action(ctx *gin.Context) {

	bytePayload, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid payload"})
		return
	}

	var payload dto.PluginActionSchema
	if err := json.Unmarshal(bytePayload, &payload); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid payload structure"})
		return
	}

	err = h.PluginService.ActionValidatePayload(payload)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid payload validation"})
		return
	}

	err = h.PluginService.Action(payload)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "success"})
}

func (h *PluginHandler) Fetch(ctx *gin.Context) {
	// TODO: validation
	// TODO: logic

	ctx.JSON(200, gin.H{})
}

func (h *PluginHandler) Get(ctx *gin.Context) {

	// TODO: validation

	res, err := h.PluginService.Get()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	ctx.JSON(200, gin.H{"plugin": res})
}
