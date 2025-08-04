package services

import (
	"github.com/google/uuid"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/repository"
	"github.com/mujak27/gamen/src/core/internal/models"
)

type ConfigurationService struct {
	repo interfaces.IConfigurationRepository
}

func NewConfigurationService(repo interfaces.IConfigurationRepository) *ConfigurationService {
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

func (s *ConfigurationService) CreateConfiguration(configuration models.Configuration) (models.Configuration, error) {
	// validate configuration based on its type
	configurationType, err := s.repo.GetConfigurationTypeServiceById(configuration.ConfigurationTypeID)
	if err != nil {
		return models.Configuration{}, err
	}
	err = configurationType.ValidateConfiguration(configuration)
	if err != nil {
		return models.Configuration{}, err
	}

	// generate new id for configuration
	configuration.ID = uuid.New()

	// create configuration and write to database
	createdConfiguration, err := s.repo.CreateConfiguration(configuration)
	if err != nil {
		return models.Configuration{}, err
	}

	// TODO: return configuration
	return createdConfiguration, nil
}

func (s *ConfigurationService) DeleteConfiguration(name string) error {
	return s.repo.DeleteConfiguration(name)
}
