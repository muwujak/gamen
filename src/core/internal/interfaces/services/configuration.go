package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/api/dto"
	"github.com/mujak27/gamen/src/core/internal/models"
)

type IConfigurationService interface {
	GetConfigurationById(ctx context.Context, payload dto.ConfigurationGetPayload) (models.Configuration, error)
	ListConfigurations(ctx context.Context, payload dto.ConfigurationListPayload) ([]models.Configuration, error)
	CreateConfiguration(ctx context.Context, payload dto.ConfigurationCreatePayload) (models.Configuration, error)
	UpdateConfiguration(ctx context.Context, payload dto.ConfigurationUpdatePayload) (models.Configuration, error)
	DeleteConfiguration(ctx context.Context, payload dto.ConfigurationDeletePayload) error
	ListConfigurationTypes() ([]models.ConfigurationType, error)
	GetConfigurationTypeById(id uuid.UUID) (models.ConfigurationType, error)
}

type IConfigurationTypeService interface {
	GetId() uuid.UUID
	Get() models.ConfigurationType
	ValidateConfiguration(configuration models.Configuration) error
}
