package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/api/dto"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/services"
)

type ConfigurationHandler struct {
	service interfaces.IConfigurationService
}

func NewConfigurationHandler(service interfaces.IConfigurationService) *ConfigurationHandler {
	return &ConfigurationHandler{service: service}
}

// TODO: how to standardize response key?
// TODO: standardize uuid checking

func (configurationHandler *ConfigurationHandler) GetConfigurationById(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid configuration ID"})
		return
	}
	configuration, err := configurationHandler.service.GetConfigurationById(uuid)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve configuration"})
		return
	}
	ctx.JSON(200, gin.H{"configuration": configuration})
}

func (configurationHandler *ConfigurationHandler) GetConfigurationTypeById(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid configuration type ID"})
		return
	}
	configurationType, err := configurationHandler.service.GetConfigurationTypeById(uuid)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve configuration type"})
		return
	}
	ctx.JSON(200, gin.H{"configuration_types": configurationType})
}

// TODO: still not properly thought, configurations logic are hard coded, but we have a function that list configurations from repository
func (configurationHandler *ConfigurationHandler) ListConfigurations(ctx *gin.Context) {
	configurationType, err := configurationHandler.service.ListConfigurations()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve configuration"})
		return
	}
	ctx.JSON(200, gin.H{"configurations": configurationType})
}

func (configurationHandler *ConfigurationHandler) ListConfigurationTypes(ctx *gin.Context) {
	configurationType, err := configurationHandler.service.ListConfigurationTypes()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve configuration type"})
		return
	}
	ctx.JSON(200, gin.H{"configuration_types": configurationType})
}

func (configurationHandler *ConfigurationHandler) CreateConfiguration(ctx *gin.Context) {
	var payload dto.ConfigurationCreatePayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	configuration, err := configurationHandler.service.CreateConfiguration(payload)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create configuration"})
		return
	}
	ctx.JSON(200, gin.H{"configuration": configuration})
}

func (configurationHandler *ConfigurationHandler) UpdateConfiguration(ctx *gin.Context) {
	var payload dto.ConfigurationUpdatePayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	configuration, err := configurationHandler.service.UpdateConfiguration(payload)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update configuration"})
		return
	}
	ctx.JSON(200, gin.H{"configuration": configuration})
}

func (configurationHandler *ConfigurationHandler) DeleteConfiguration(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid configuration ID"})
		return
	}
	err = configurationHandler.service.DeleteConfiguration(uuid)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete configuration"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Configuration deleted successfully"})
}
