package services

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/api/dto"
	repositoryMocks "github.com/mujak27/gamen/src/core/internal/interfaces/repository/mocks"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/services"
	"github.com/mujak27/gamen/src/core/internal/interfaces/services/mocks"
	"github.com/mujak27/gamen/src/core/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestWidgetServiceTypeAssertion(t *testing.T) {
	var _ interfaces.WidgetService = &WidgetService{}
}

// ACTION TEST

func TestWidgetServiceActionSuccess(t *testing.T) {
	mockCatalogueService := mocks.NewMockCatalogueService(t)
	mockConfigurationService := mocks.NewMockConfigurationService(t)
	mockWidgetRepository := repositoryMocks.NewMockWidgetRepository(t)
	mockPluginService := mocks.NewMockPluginService(t)

	configurationUUID := uuid.New()
	pluginUUID := uuid.New()
	widgetUUID := uuid.New()

	configuration := models.Configuration{
		ID: configurationUUID,
		Data: map[string]interface{}{
			"key": "value",
		},
	}

	payload := dto.WidgetActionPayload{
		WidgetID: widgetUUID,
		Data:     nil,
	}

	mockWidgetRepository.On("GetWidgetById", payload.WidgetID).Return(models.Widget{
		ID:       widgetUUID,
		PluginID: pluginUUID,
		Plugin: models.Plugin{
			ID:          pluginUUID,
			Name:        "test",
			Description: "test",
			Version:     "1.0.0",
		},
		ConfigurationID: configurationUUID,
		Configuration:   configuration,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}, nil)

	mockCatalogueService.On("GetPluginFunctionById", pluginUUID.String()).Return(mockPluginService, nil)

	mockPluginService.On("Action", dto.PluginActionSchema{
		Configuration: configuration,
		RawPluginData: payload.Data,
	}).Return(nil)

	widgetService := NewWidgetService(
		mockCatalogueService,
		mockConfigurationService,
		mockWidgetRepository,
	)

	err := widgetService.Action(payload)
	assert.NoError(t, err)

}
