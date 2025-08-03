package interfaces

import (
	"github.com/google/uuid"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/services"
)

type CatalogueRepository interface {
	GetPluginFunctionById(id uuid.UUID) (interfaces.PluginService, error)
	RegisterPluginFunction(key uuid.UUID, plugin interfaces.PluginService)
	ListPluginKeys() []uuid.UUID
}
