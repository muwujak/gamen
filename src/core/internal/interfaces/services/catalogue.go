package interfaces

import (
	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/models"
)

type CatalogueService interface {
	GetPluginModelById(id uuid.UUID) (models.Plugin, error)
	GetPluginFunctionById(id uuid.UUID) (PluginService, error)
	ListPluginKeys() ([]uuid.UUID, error)
}
