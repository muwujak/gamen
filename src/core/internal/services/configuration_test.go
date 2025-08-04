package services

import (
	"testing"

	"github.com/google/uuid"
	kubernetesConfigurationType "github.com/mujak27/gamen/src/core/internal/extensions/configuration_types/kubernetes/services"
	"github.com/mujak27/gamen/src/core/internal/interfaces/repository/mocks"
	serviceInterface "github.com/mujak27/gamen/src/core/internal/interfaces/services"
	"github.com/mujak27/gamen/src/core/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestConfigurationServiceTypeAssertion(t *testing.T) {
	var _ serviceInterface.IConfigurationService = &ConfigurationService{}
}

func TestConfigurationServiceCreateConfiguration(t *testing.T) {

	configuration := models.Configuration{
		ConfigurationTypeID: uuid.New(),
		Data: map[string]interface{}{
			"api-server-endpoint": "https://kubernetes.default.svc",
			"token":               "token",
		},
	}

	// use real logic from kubernetes configuration type service
	kubernetesConfigurationTypeService := kubernetesConfigurationType.NewKubernetesConfigurationTypeService()

	// mock configuration repository to return kubernetes configuration type service
	configurationRepo := mocks.NewMockIConfigurationRepository(t)
	configurationRepo.EXPECT().GetConfigurationTypeServiceById(mock.Anything).Return(kubernetesConfigurationTypeService, nil).Twice()
	configurationRepo.EXPECT().CreateConfiguration(mock.Anything).Return(configuration, nil)

	configurationService := NewConfigurationService(configurationRepo)

	_, err := configurationService.CreateConfiguration(configuration)
	assert.NoError(t, err)

	// test invalid configuration, should error from kubernetes configuration type service validation
	configuration2 := models.Configuration{
		ConfigurationTypeID: uuid.New(),
		Data: map[string]interface{}{
			"invalid-key": "https://kubernetes.default.svc",
			"token":       "token",
		},
	}

	_, err = configurationService.CreateConfiguration(configuration2)
	assert.Error(t, err)
}
