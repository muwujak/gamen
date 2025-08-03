package interfaces

import (
	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/api/dto"
	"github.com/mujak27/gamen/src/core/internal/models"
)

type PluginService interface { // interface for each plugin
	Action(dto.PluginActionSchema) error
	ActionValidatePayload(dto.PluginActionSchema) error
	Fetch(string) (dto.PluginFetchSchema, error) // each catalogue should implement its own model validation
	Get() (models.Plugin, error)
	GetId() uuid.UUID // needed when registering routes
}
