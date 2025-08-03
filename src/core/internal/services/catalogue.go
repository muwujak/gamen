package services

import (
	"github.com/google/uuid"
	repositoryInterface "github.com/mujak27/gamen/src/core/internal/interfaces/repository"
	serviceInterface "github.com/mujak27/gamen/src/core/internal/interfaces/services"
	"github.com/mujak27/gamen/src/core/internal/models"
)

type CatalogueService struct {
	repo                 repositoryInterface.CatalogueRepository
	configurationService *ConfigurationService
}

func NewCatalogueService(repo repositoryInterface.CatalogueRepository, configurationService *ConfigurationService) *CatalogueService {
	return &CatalogueService{
		repo:                 repo,
		configurationService: configurationService,
	}
}

func (s *CatalogueService) GetPluginModelById(id uuid.UUID) (models.Plugin, error) {
	plugin, err := s.repo.GetPluginFunctionById(id)
	if err != nil {
		return models.Plugin{}, err
	}
	model, err := plugin.Get()
	if err != nil {
		return models.Plugin{}, err
	}
	return model, nil
}

func (s *CatalogueService) GetPluginFunctionById(id uuid.UUID) (serviceInterface.PluginService, error) {
	plugin, err := s.repo.GetPluginFunctionById(id)
	if err != nil {
		return nil, err
	}
	return plugin, nil
}

func (s *CatalogueService) ListPluginKeys() ([]uuid.UUID, error) {
	plugins := s.repo.ListPluginKeys()
	return plugins, nil
}
