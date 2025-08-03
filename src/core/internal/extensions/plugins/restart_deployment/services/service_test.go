package services

import (
	"testing"

	"github.com/mujak27/gamen/src/core/internal/api/dto"
	kubernetes_model "github.com/mujak27/gamen/src/core/internal/extensions/configuration_types/kubernetes/dto"
	kubernetes_configuration_type "github.com/mujak27/gamen/src/core/internal/extensions/configuration_types/kubernetes/interfaces/mocks"
	pluginMocks "github.com/mujak27/gamen/src/core/internal/extensions/plugins/restart_deployment/interfaces/mocks"
	"github.com/mujak27/gamen/src/core/internal/interfaces/repository/mocks"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/services"
	"github.com/mujak27/gamen/src/core/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRestartDeploymentPluginServiceTypeAssertion(t *testing.T) {
	var _ interfaces.PluginService = &RestartDeploymentPluginService{}
}

func TestRestartDeploymentPluginServiceMethodAction(t *testing.T) {

	kubernetesConfiguration := kubernetes_model.KubernetesConfiguration{
		ApiServerEndpoint: "https://dummy-kubernetes-endpoint",
		Token:             "token",
	}
	var data models.JSON
	err := data.Transform(kubernetesConfiguration)
	if err != nil {
		t.Fatalf("Error scanning kubernetes configuration: %v", err)
	}

	configuration := models.Configuration{
		Data: data,
	}

	payload := dto.PluginActionSchema{
		Configuration: configuration,
		RawPluginData: map[string]interface{}{
			"deployment_name": "test",
		},
	}

	// configurationTypeID := uuidUtil.GenerateSHA1UUID("configuration-type-kubernetes")
	mockKubernetesConfigurationService := kubernetes_configuration_type.NewMockIKubernetesService(t)
	mockKubernetesConfigurationService.EXPECT().CreateKubernetesClient(mock.Anything).Return(nil, nil).Once()
	mockCatalogueRepository := mocks.NewMockCatalogueRepository(t)
	mockUtilService := pluginMocks.NewMockIUtilService(t)
	mockUtilService.EXPECT().RestartDeployment(mock.Anything, payload.RawPluginData["deployment_name"]).Return(nil).Once()
	restartDeploymentService := NewRestartDeploymentService(mockKubernetesConfigurationService, mockCatalogueRepository, mockUtilService)
	assert.NoError(t, restartDeploymentService.Action(payload))
}
