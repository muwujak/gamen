package interfaces

import (
	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/api/dto"
	"github.com/mujak27/gamen/src/core/internal/models"
)

type IConfigurationService interface {
	GetConfigurationById(id uuid.UUID) (models.Configuration, error)
	ListConfigurations() ([]models.Configuration, error)
	CreateConfiguration(payload dto.ConfigurationCreatePayload) (models.Configuration, error)
	UpdateConfiguration(payload dto.ConfigurationUpdatePayload) (models.Configuration, error)
	DeleteConfiguration(id uuid.UUID) error
	ListConfigurationTypes() ([]models.ConfigurationType, error)
	GetConfigurationTypeById(id uuid.UUID) (models.ConfigurationType, error)
}

type IConfigurationTypeService interface {
	GetId() uuid.UUID
	Get() models.ConfigurationType
	ValidateConfiguration(configuration models.Configuration) error
}
