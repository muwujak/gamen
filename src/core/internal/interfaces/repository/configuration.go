package interfaces

import (
	"github.com/mujak27/gamen/src/core/internal/models"
)

type ConfigurationRepository interface {
	GetConfigurationById(id string) (models.Configuration, error)
	GetConfigurationTypeById(id string) (models.ConfigurationType, error)
	ListConfigurations() ([]models.Configuration, error)
	ListConfigurationTypes() ([]models.ConfigurationType, error)
	CreateConfiguration(configuration models.Configuration) (models.Configuration, error)
	DeleteConfiguration(name string) error
}
