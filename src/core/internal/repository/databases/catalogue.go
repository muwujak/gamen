package databases

import (
	"fmt"

	"github.com/google/uuid"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/services"
	"github.com/mujak27/gamen/src/core/internal/models"
)

type CatalogueRepository struct {
	db              interface{}
	pluginFunctions map[uuid.UUID]interfaces.PluginService
}

func NewCatalogueRepository(db interface{}) *CatalogueRepository {
	return &CatalogueRepository{db: db, pluginFunctions: make(map[uuid.UUID]interfaces.PluginService)}
}

func (r *CatalogueRepository) GetPluginFunctionById(id uuid.UUID) (interfaces.PluginService, error) {
	plugin := r.pluginFunctions[id]
	if plugin == nil {
		return nil, fmt.Errorf("plugin not found")
	}
	return plugin, nil
}

func (r *CatalogueRepository) GetPluginModelById(id uuid.UUID) (models.Plugin, error) {
	plugin, err := r.GetPluginFunctionById(id)
	if err != nil {
		return models.Plugin{}, err
	}
	model, err := plugin.Get()
	if err != nil {
		return models.Plugin{}, err
	}
	return model, nil
}

func (r *CatalogueRepository) RegisterPluginFunction(key uuid.UUID, plugin interfaces.PluginService) {
	// check if key already exists
	if _, exists := r.pluginFunctions[key]; exists {
		panic(fmt.Sprintf("plugin function with key %s already exists", key.String()))
	}
	r.pluginFunctions[key] = plugin
}

func (r *CatalogueRepository) ListPluginKeys() []uuid.UUID {
	keys := make([]uuid.UUID, 0, len(r.pluginFunctions))
	for k := range r.pluginFunctions {
		keys = append(keys, k)
	}
	return keys
}
