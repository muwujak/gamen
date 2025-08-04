package services

import (
	"fmt"
	"testing"

	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/services"
	"github.com/mujak27/gamen/src/core/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestKubernetesConfigurationServiceTypeAssertion(t *testing.T) {
	var _ interfaces.IConfigurationTypeService = &KubernetesConfigurationTypeService{}
}

func TestKubernetesConfigurationTypeServiceValidateConfiguration(t *testing.T) {
	service := NewKubernetesConfigurationTypeService()

	configuration1 := models.Configuration{
		Data: map[string]interface{}{
			"api-server-endpoint": "https://kubernetes.default.svc",
			"token":               "token",
		},
	}

	err := service.ValidateConfiguration(configuration1)
	assert.NoError(t, err)

	// invalid key
	configuration2 := models.Configuration{
		Data: map[string]interface{}{
			"invalid-endpoint": "https://kubernetes.default.svc",
			"token":            "token",
		},
	}

	err = service.ValidateConfiguration(configuration2)
	fmt.Println("err", err)
	assert.Error(t, err)

	// invalid url
	configuration3 := models.Configuration{
		Data: map[string]interface{}{
			"api-server-endpoint": "invalid-url",
			"token":               "token",
		},
	}

	err = service.ValidateConfiguration(configuration3)
	fmt.Println("err", err)
	assert.Error(t, err)
}
