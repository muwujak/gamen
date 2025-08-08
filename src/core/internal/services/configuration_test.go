package services

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/api/dto"
	"github.com/mujak27/gamen/src/core/internal/api/middleware"
	kubernetesConfigurationType "github.com/mujak27/gamen/src/core/internal/extensions/configuration_types/kubernetes/services"
	"github.com/mujak27/gamen/src/core/internal/interfaces/repository/mocks"
	serviceInterface "github.com/mujak27/gamen/src/core/internal/interfaces/services"
	serviceMocks "github.com/mujak27/gamen/src/core/internal/interfaces/services/mocks"
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
	mockConfigurationRepo := mocks.NewMockIConfigurationRepository(t)
	mockUserService := serviceMocks.NewMockIUserService(t)

	configurationService := NewConfigurationService(mockConfigurationRepo, mockUserService)

	// CASE: create configuration without user id in context
	blankContext := context.Background()
	_, err := configurationService.CreateConfiguration(blankContext, validPayload)
	assert.Error(t, err)

	// CASE: create valid configuration with user id in context
	mockConfigurationRepo.EXPECT().GetConfigurationTypeServiceById(mock.Anything).Return(kubernetesConfigurationTypeService, nil).Once()
	mockConfigurationRepo.EXPECT().CreateConfiguration(mock.Anything).Return(models.Configuration{}, nil)
	mockUserService.EXPECT().IsTeamMember(mock.Anything, mock.Anything).Return(true, nil).Once()
	validContext := context.WithValue(blankContext, middleware.UserIDKey, uuid.New().String())
	_, err = configurationService.CreateConfiguration(validContext, validPayload)
	assert.NoError(t, err)

	// CASE: create invalid configuration with user id in context
	_, err = configurationService.CreateConfiguration(validContext, invalidPayload)
	assert.Error(t, err)
}

// TODO: split test into multiple based on test cases
func TestConfigurationServiceConfigurationLifecycle(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("mock.db"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	seeding.Migrate(db)

	// use real logic from kubernetes configuration type service
	kubernetesConfigurationTypeService := kubernetesConfigurationType.NewKubernetesConfigurationTypeService()

	// TODO: use valid teamID
	teamId := uuid.New()
	userId := uuid.New()
	validPayload := dto.ConfigurationCreatePayload{
		Name:                "homelab-kubernetes",
		Description:         "testing purpose",
		TeamID:              teamId,
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

	userService := serviceMocks.NewMockIUserService(t)

	configurationService := NewConfigurationService(configurationRepo, userService)

	// CASE: create configuration without user id in context
	blankContext := context.Background()
	_, err = configurationService.CreateConfiguration(blankContext, validPayload)
	assert.Error(t, err)

	// CASE: create valid configuration with user id in context
	userService.EXPECT().IsTeamMember(userId, teamId).Return(true, nil).Once()
	validContext := context.WithValue(blankContext, middleware.UserIDKey, userId.String())

	createdConfiguration, err := configurationService.CreateConfiguration(validContext, validPayload)
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

	// CASE: update configuration with matching user id in context
	userService.EXPECT().IsTeamMember(userId, teamId).Return(true, nil).Once()
	updatedConfiguration, err := configurationService.UpdateConfiguration(validContext, updatedPayload)
	assert.NoError(t, err)
	assert.NotNil(t, updatedConfiguration)
	assert.Equal(t, updatedConfiguration.Name, "homelab-kubernetes-updated")
	assert.Equal(t, updatedConfiguration.Description, "testing purpose updated")

	// CASE: delete configuration with matching user id in context
	userService.EXPECT().IsTeamMember(userId, teamId).Return(true, nil).Once()
	deletePayload := dto.ConfigurationDeletePayload{
		ID: createdConfiguration.ID,
	}
	err = configurationService.DeleteConfiguration(validContext, deletePayload)
	assert.NoError(t, err)
}
