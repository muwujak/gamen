package services

import (
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/repository"
	"github.com/mujak27/gamen/src/core/internal/models"
)

type ConfigurationService struct {
	repo interfaces.ConfigurationRepository
}

func NewConfigurationService(repo interfaces.ConfigurationRepository) *ConfigurationService {
	return &ConfigurationService{repo: repo}
}

func (s *ConfigurationService) GetConfigurationById(id string) (models.Configuration, error) {
	return s.repo.GetConfigurationById(id)
}

func (s *ConfigurationService) GetConfigurationTypeById(id string) (models.ConfigurationType, error) {
	return s.repo.GetConfigurationTypeById(id)
}

func (s *ConfigurationService) ListConfigurations() ([]models.Configuration, error) {
	return s.repo.ListConfigurations()
}

func (s *ConfigurationService) ListConfigurationTypes() ([]models.ConfigurationType, error) {
	return s.repo.ListConfigurationTypes()
}

// TODO: create repository function for create configuration to database
func (s *ConfigurationService) CreateConfiguration(configuration models.Configuration) (models.Configuration, error) {
	return s.repo.CreateConfiguration(configuration)
}

func (s *ConfigurationService) DeleteConfiguration(name string) error {
	return s.repo.DeleteConfiguration(name)
}
