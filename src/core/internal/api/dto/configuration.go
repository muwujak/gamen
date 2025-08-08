package dto

import (
	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/models"
)

type ConfigurationCreatePayload struct {
	Name                string      `json:"name" validate:"required"`
	Description         string      `json:"description" validate:"required"`
	ConfigurationTypeID uuid.UUID   `json:"configuration_type_id" validate:"required"`
	TeamID              uuid.UUID   `json:"team_id" validate:"required"`
	Data                models.JSON `json:"data" validate:"required"`
}

type ConfigurationUpdatePayload struct {
	ID          uuid.UUID   `json:"id" validate:"required"`
	Name        string      `json:"name" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Data        models.JSON `json:"data" validate:"required"`
}

type ConfigurationListPayload struct {
	TeamID uuid.UUID `json:"team_id" validate:"required"`
}

type ConfigurationGetPayload struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type ConfigurationDeletePayload struct {
	ID uuid.UUID `json:"id" validate:"required"`
}
