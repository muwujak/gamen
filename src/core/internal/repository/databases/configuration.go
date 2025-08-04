package databases

import (
	"fmt"

	"github.com/google/uuid"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/services"
	"github.com/mujak27/gamen/src/core/internal/models"
	"gorm.io/gorm"
)

type ConfigurationRepository struct {
	db                        *gorm.DB
	configurationTypeServices map[uuid.UUID]interfaces.IConfigurationTypeService
}

func NewConfigurationRepository(db *gorm.DB) *ConfigurationRepository {
	return &ConfigurationRepository{
		db:                        db,
		configurationTypeServices: make(map[uuid.UUID]interfaces.IConfigurationTypeService),
	}
}

func (r *ConfigurationRepository) GetConfigurationById(name string) (models.Configuration, error) {
	var configuration models.Configuration
	err := r.db.First(&configuration, "name = ?", name).Error
	if err != nil {
		return models.Configuration{}, err
	}
	return configuration, nil
}

func (r *ConfigurationRepository) ListConfigurations() ([]models.Configuration, error) {
	var configurations []models.Configuration
	err := r.db.Find(&configurations).Error
	if err != nil {
		return []models.Configuration{}, err
	}
	return configurations, nil
}

func (r *ConfigurationRepository) CreateConfiguration(configuration models.Configuration) (models.Configuration, error) {
	err := r.db.Create(&configuration).Error
	if err != nil {
		return models.Configuration{}, err
	}
	return configuration, nil
}

func (r *ConfigurationRepository) DeleteConfiguration(name string) error {
	err := r.db.Delete(&models.Configuration{}, "name = ?", name).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ConfigurationRepository) RegisterConfigurationTypeService(key uuid.UUID, plugin interfaces.IConfigurationTypeService) error {
	// check if key already exists
	if _, exists := r.configurationTypeServices[key]; exists {
		return fmt.Errorf("plugin function with key %s already exists", key.String())
	}
	r.configurationTypeServices[key] = plugin
	return nil
}

func (r *ConfigurationRepository) GetConfigurationTypeServiceById(id uuid.UUID) (interfaces.IConfigurationTypeService, error) {
	configurationTypeService := r.configurationTypeServices[id]
	if configurationTypeService == nil {
		return nil, fmt.Errorf("configuration type service not found")
	}
	return configurationTypeService, nil
}

func (r *ConfigurationRepository) GetConfigurationTypeById(name string) (models.ConfigurationType, error) {
	var configurationType models.ConfigurationType
	err := r.db.First(&configurationType, "name = ?", name).Error
	if err != nil {
		return models.ConfigurationType{}, err
	}
	return configurationType, nil
}

func (r *ConfigurationRepository) ListConfigurationTypes() ([]models.ConfigurationType, error) {
	var configurationTypes []models.ConfigurationType
	err := r.db.Find(&configurationTypes).Error
	if err != nil {
		return []models.ConfigurationType{}, err
	}
	return configurationTypes, nil
}
