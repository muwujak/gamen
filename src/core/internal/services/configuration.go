package services

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/api/dto"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/repository"
	"github.com/mujak27/gamen/src/core/internal/models"
)

type ConfigurationService struct {
	repo interfaces.IConfigurationRepository
}

func NewConfigurationService(repo interfaces.IConfigurationRepository) *ConfigurationService {
	return &ConfigurationService{repo: repo}
}

func (s *ConfigurationService) GetConfigurationById(id uuid.UUID) (models.Configuration, error) {
	return s.repo.GetConfigurationById(id)
}

func (s *ConfigurationService) GetConfigurationTypeById(id uuid.UUID) (models.ConfigurationType, error) {
	return s.repo.GetConfigurationTypeById(id)
}

func (s *ConfigurationService) ListConfigurations() ([]models.Configuration, error) {
	return s.repo.ListConfigurations()
}

func (s *ConfigurationService) ListConfigurationTypes() ([]models.ConfigurationType, error) {
	return s.repo.ListConfigurationTypes()
}

func (s *ConfigurationService) CreateConfiguration(payload dto.ConfigurationCreatePayload) (models.Configuration, error) {

	// validate payload bindings
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return models.Configuration{}, err
	}

	// TODO: need to get current user id from context
	configuration := models.Configuration{
		ID:                  uuid.New(),
		Name:                payload.Name,
		Description:         payload.Description,
		ConfigurationTypeID: payload.ConfigurationTypeID,
		TeamID:              payload.TeamID,
		Data:                payload.Data,
		CreatedBy:           uuid.New(),
		CreatedAt:           time.Now().UTC(),
		UpdatedAt:           time.Now().UTC(),
		IsActive:            true,
		LastUsedAt:          nil,
	}

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

	return createdConfiguration, nil
}

func (s *ConfigurationService) UpdateConfiguration(payload dto.ConfigurationUpdatePayload) (models.Configuration, error) {

	// validate payload bindings
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return models.Configuration{}, err
	}

	// get existing configuration
	existingConfiguration, err := s.repo.GetConfigurationById(payload.ID)
	if err != nil {
		return models.Configuration{}, err
	}

	// validate configuration based on its type
	configurationType, err := s.repo.GetConfigurationTypeServiceById(existingConfiguration.ConfigurationTypeID)
	if err != nil {
		return models.Configuration{}, err
	}

	// update configuration
	existingConfiguration.Name = payload.Name
	existingConfiguration.Description = payload.Description
	existingConfiguration.Data = payload.Data
	existingConfiguration.UpdatedAt = time.Now().UTC()

	err = configurationType.ValidateConfiguration(existingConfiguration)
	if err != nil {
		return models.Configuration{}, err
	}

	// update configuration and write to database
	updatedConfiguration, err := s.repo.UpdateConfiguration(existingConfiguration)
	if err != nil {
		return models.Configuration{}, err
	}

	return updatedConfiguration, nil
}

func (s *ConfigurationService) DeleteConfiguration(id uuid.UUID) error {

	// TODO: authorization check

	return s.repo.DeleteConfiguration(id)
}
