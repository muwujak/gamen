package services

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/api/dto"
	kubernetesConfigurationType "github.com/mujak27/gamen/src/core/internal/extensions/configuration_types/kubernetes/services"
	"github.com/mujak27/gamen/src/core/internal/interfaces/repository/mocks"
	serviceInterface "github.com/mujak27/gamen/src/core/internal/interfaces/services"
	"github.com/mujak27/gamen/src/core/internal/models"
	"github.com/mujak27/gamen/src/core/internal/repository/databases"
	"github.com/mujak27/gamen/src/core/internal/seeding"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestConfigurationServiceTypeAssertion(t *testing.T) {
	var _ serviceInterface.IConfigurationService = &ConfigurationService{}
}

func TestConfigurationServiceCreateConfiguration(t *testing.T) {

	// TODO: use valid teamID
	validPayload := dto.ConfigurationCreatePayload{
		Name:                "test-configuration",
		Description:         "test-description",
		TeamID:              uuid.New(),
		ConfigurationTypeID: uuid.New(),
		Data: map[string]interface{}{
			"api-server-endpoint": "https://kubernetes.default.svc",
			"token":               "token",
		},
	}

	// test invalid configuration, should error from kubernetes configuration type service validation
	invalidPayload := dto.ConfigurationCreatePayload{
		ConfigurationTypeID: uuid.New(),
		Data: map[string]interface{}{
			"invalid-key": "https://kubernetes.default.svc",
			"token":       "token",
		},
	}

	// use real logic from kubernetes configuration type service
	kubernetesConfigurationTypeService := kubernetesConfigurationType.NewKubernetesConfigurationTypeService()

	// mock configuration repository to return kubernetes configuration type service
	configurationRepo := mocks.NewMockIConfigurationRepository(t)
	configurationRepo.EXPECT().GetConfigurationTypeServiceById(mock.Anything).Return(kubernetesConfigurationTypeService, nil).Once()
	configurationRepo.EXPECT().CreateConfiguration(mock.Anything).Return(models.Configuration{}, nil)

	configurationService := NewConfigurationService(configurationRepo)

	_, err := configurationService.CreateConfiguration(validPayload)
	assert.NoError(t, err)

	_, err = configurationService.CreateConfiguration(invalidPayload)
	assert.Error(t, err)
}

func TestConfigurationServiceCreateAndDeleteConfiguration(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("mock.db"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	seeding.Migrate(db)

	// use real logic from kubernetes configuration type service
	kubernetesConfigurationTypeService := kubernetesConfigurationType.NewKubernetesConfigurationTypeService()

	// TODO: use valid teamID
	validPayload := dto.ConfigurationCreatePayload{
		Name:                "homelab-kubernetes",
		Description:         "testing purpose",
		TeamID:              uuid.New(),
		ConfigurationTypeID: kubernetesConfigurationTypeService.GetId(),
		Data: map[string]interface{}{
			"api-server-endpoint": "https://kubernetes.default.svc",
			"token":               "token",
		},
	}

	// mock configuration repository to return kubernetes configuration type service
	configurationRepo := databases.NewConfigurationRepository(db)
	fmt.Println("kubernetesConfigurationTypeService.GetId(): ", kubernetesConfigurationTypeService.GetId())
	configurationRepo.RegisterConfigurationTypeService(kubernetesConfigurationTypeService.GetId(), kubernetesConfigurationTypeService)

	configurationService := NewConfigurationService(configurationRepo)

	createdConfiguration, err := configurationService.CreateConfiguration(validPayload)
	fmt.Println("error: ", err)
	fmt.Println(createdConfiguration)
	assert.NoError(t, err)
	assert.NotNil(t, createdConfiguration)

	updatedPayload := dto.ConfigurationUpdatePayload{
		ID:          createdConfiguration.ID,
		Name:        "homelab-kubernetes-updated",
		Description: "testing purpose updated",
		Data: map[string]interface{}{
			"api-server-endpoint": "https://kubernetes.default.svc",
			"token":               "token-updated",
		},
	}

	updatedConfiguration, err := configurationService.UpdateConfiguration(updatedPayload)
	fmt.Println("error: ", err)
	fmt.Println(updatedConfiguration)
	assert.NoError(t, err)
	assert.NotNil(t, updatedConfiguration)
	assert.Equal(t, updatedConfiguration.Name, "homelab-kubernetes-updated")
	assert.Equal(t, updatedConfiguration.Description, "testing purpose updated")

	// TODO: use delete service instead of direct db delete
	err = configurationService.DeleteConfiguration(createdConfiguration.ID)
	assert.NoError(t, err)
}
