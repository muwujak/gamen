package interfaces

import (
	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/models"
)

type IConfigurationService interface {
	GetConfigurationById(id string) (models.Configuration, error)
	GetConfigurationTypeById(id string) (models.ConfigurationType, error)
	ListConfigurations() ([]models.Configuration, error)
	ListConfigurationTypes() ([]models.ConfigurationType, error)
	CreateConfiguration(models.Configuration) (models.Configuration, error)
	DeleteConfiguration(name string) error
}

type IConfigurationTypeService interface {
	GetId() uuid.UUID
	Get() models.ConfigurationType
	ValidateConfiguration(configuration models.Configuration) error
}
