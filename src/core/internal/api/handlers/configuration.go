package handlers

import (
	"github.com/gin-gonic/gin"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/services"
	"github.com/mujak27/gamen/src/core/internal/models"
)

type ConfigurationHandler struct {
	service interfaces.IConfigurationService
}

func NewConfigurationHandler(service interfaces.IConfigurationService) *ConfigurationHandler {
	return &ConfigurationHandler{service: service}
}

// TODO: how to standardize response key?

func (configurationHandler *ConfigurationHandler) GetConfigurationById(ctx *gin.Context) {
	id := ctx.Param("id")
	configuration, err := configurationHandler.service.GetConfigurationById(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve configuration"})
		return
	}
	ctx.JSON(200, gin.H{"configuration": configuration})
}

func (configurationHandler *ConfigurationHandler) GetConfigurationTypeById(ctx *gin.Context) {
	id := ctx.Param("id")
	configurationType, err := configurationHandler.service.GetConfigurationTypeById(id)
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
	var configuration models.Configuration
	if err := ctx.ShouldBindJSON(&configuration); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := configurationHandler.service.CreateConfiguration(configuration)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve configuration"})
		return
	}
	ctx.JSON(200, gin.H{"configuration": configuration})
}

func (configurationHandler *ConfigurationHandler) DeleteConfigurationType(ctx *gin.Context) {
	name := ctx.Param("id")
	err := configurationHandler.service.DeleteConfiguration(name)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete configuration type"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Configuration type deleted successfully"})
}
