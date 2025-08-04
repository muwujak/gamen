package interfaces

import (
	"github.com/google/uuid"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/services"
	"github.com/mujak27/gamen/src/core/internal/models"
)

type IConfigurationRepository interface {
	GetConfigurationById(id string) (models.Configuration, error)
	GetConfigurationTypeServiceById(uuid.UUID) (interfaces.IConfigurationTypeService, error)
	GetConfigurationTypeById(id string) (models.ConfigurationType, error)
	ListConfigurations() ([]models.Configuration, error)
	ListConfigurationTypes() ([]models.ConfigurationType, error)
	CreateConfiguration(configuration models.Configuration) (models.Configuration, error)
	DeleteConfiguration(name string) error
}
