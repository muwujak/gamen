package services

import (
	"context"
	"errors"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/api/dto"
	"github.com/mujak27/gamen/src/core/internal/api/middleware"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/repository"
	serviceInterfaces "github.com/mujak27/gamen/src/core/internal/interfaces/services"
	"github.com/mujak27/gamen/src/core/internal/models"
)

type ConfigurationService struct {
	repo        interfaces.IConfigurationRepository
	userService serviceInterfaces.IUserService
}

func NewConfigurationService(repo interfaces.IConfigurationRepository, userService serviceInterfaces.IUserService) *ConfigurationService {
	return &ConfigurationService{repo: repo, userService: userService}
}

func (s *ConfigurationService) GetUserId(ctx context.Context) (uuid.UUID, error) {
	// get user id from context
	userIdRaw := ctx.Value(middleware.UserIDKey)
	if userIdRaw == nil {
		return uuid.Nil, errors.New("user id not found in context")
	}
	userIdRawStr := userIdRaw.(string)
	userId, err := uuid.Parse(userIdRawStr)
	if err != nil {
		return uuid.Nil, err
	}
	return userId, nil
}

func (s *ConfigurationService) GetConfigurationById(ctx context.Context, payload dto.ConfigurationGetPayload) (models.Configuration, error) {
	userId, err := s.GetUserId(ctx)
	if err != nil {
		return models.Configuration{}, err
	}
	configuration, err := s.repo.GetConfigurationById(payload.ID)
	if err != nil {
		return models.Configuration{}, err
	}
	isTeamMember, err := s.userService.IsTeamMember(userId, configuration.TeamID)
	if err != nil {
		return models.Configuration{}, err
	}
	if !isTeamMember {
		return models.Configuration{}, errors.New("user does not have access to the configuration")
	}
	return configuration, nil
}

func (s *ConfigurationService) ListConfigurations(ctx context.Context, payload dto.ConfigurationListPayload) ([]models.Configuration, error) {
	userId, err := s.GetUserId(ctx)
	if err != nil {
		return []models.Configuration{}, err
	}
	isTeamMember, err := s.userService.IsTeamMember(userId, payload.TeamID)
	if err != nil {
		return []models.Configuration{}, err
	}
	if !isTeamMember {
		return []models.Configuration{}, errors.New("user does not have access to the team")
	}
	return s.repo.ListConfigurationsByTeamId(payload.TeamID)
}

func (s *ConfigurationService) CreateConfiguration(ctx context.Context, payload dto.ConfigurationCreatePayload) (models.Configuration, error) {

	// validate payload bindings
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		return models.Configuration{}, err
	}

	userId, err := s.GetUserId(ctx)
	if err != nil {
		return models.Configuration{}, err
	}
	isTeamMember, err := s.userService.IsTeamMember(userId, payload.TeamID)
	if err != nil {
		return models.Configuration{}, err
	}
	if !isTeamMember {
		return models.Configuration{}, errors.New("user does not have access to the team")
	}

	configuration := models.Configuration{
		ID:                  uuid.New(),
		Name:                payload.Name,
		Description:         payload.Description,
		ConfigurationTypeID: payload.ConfigurationTypeID,
		TeamID:              payload.TeamID,
		Data:                payload.Data,
		CreatedBy:           userId,
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

func (s *ConfigurationService) UpdateConfiguration(ctx context.Context, payload dto.ConfigurationUpdatePayload) (models.Configuration, error) {

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

	userId, err := s.GetUserId(ctx)
	if err != nil {
		return models.Configuration{}, err
	}

	isTeamMember, err := s.userService.IsTeamMember(userId, existingConfiguration.TeamID)
	if err != nil {
		return models.Configuration{}, err
	}
	if !isTeamMember {
		return models.Configuration{}, errors.New("user does not have access to the configuration")
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

func (s *ConfigurationService) DeleteConfiguration(ctx context.Context, payload dto.ConfigurationDeletePayload) error {

	userId, err := s.GetUserId(ctx)
	if err != nil {
		return err
	}

	existingConfiguration, err := s.repo.GetConfigurationById(payload.ID)
	if err != nil {
		return err
	}

	isTeamMember, err := s.userService.IsTeamMember(userId, existingConfiguration.TeamID)
	if err != nil {
		return err
	}
	if !isTeamMember {
		return errors.New("user does not have access to the configuration")
	}

	return s.repo.DeleteConfiguration(payload.ID)
}

func (s *ConfigurationService) GetConfigurationTypeById(id uuid.UUID) (models.ConfigurationType, error) {
	return s.repo.GetConfigurationTypeById(id)
}

func (s *ConfigurationService) ListConfigurationTypes() ([]models.ConfigurationType, error) {
	return s.repo.ListConfigurationTypes()
}
