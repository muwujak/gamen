package services

import (
	"testing"

	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/services"
	"github.com/mujak27/gamen/src/core/internal/models"
)

func TestCatalogueServiceTypeAssertion(t *testing.T) {
	var _ interfaces.CatalogueService = &CatalogueService{}
}

type mockCatalogueService struct{}

func (m *mockCatalogueService) GetPluginModelById(id string) (models.Plugin, error) {
	return models.Plugin{}, nil
}

func (m *mockCatalogueService) GetPluginFunctionById(id string) (interfaces.PluginService, error) {
	return nil, nil
}

func (m *mockCatalogueService) ListPluginKeys() ([]string, error) {
	return nil, nil
}
