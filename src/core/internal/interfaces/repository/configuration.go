package interfaces

import (
	"github.com/google/uuid"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/services"
	"github.com/mujak27/gamen/src/core/internal/models"
)

type IConfigurationRepository interface {
	GetConfigurationById(id uuid.UUID) (models.Configuration, error)
	GetConfigurationTypeServiceById(uuid.UUID) (interfaces.IConfigurationTypeService, error)
	GetConfigurationTypeById(id uuid.UUID) (models.ConfigurationType, error)
	ListConfigurationsByTeamId(teamId uuid.UUID) ([]models.Configuration, error)
	ListConfigurationTypes() ([]models.ConfigurationType, error)
	CreateConfiguration(configuration models.Configuration) (models.Configuration, error)
	UpdateConfiguration(configuration models.Configuration) (models.Configuration, error)
	DeleteConfiguration(id uuid.UUID) error
}
