package databases

import (
	"github.com/mujak27/gamen/src/core/internal/models"
	"gorm.io/gorm"
)

type ConfigurationRepository struct {
	db *gorm.DB
}

func NewConfigurationRepository(db *gorm.DB) *ConfigurationRepository {
	return &ConfigurationRepository{db: db}
}

func (r *ConfigurationRepository) GetConfigurationById(name string) (models.Configuration, error) {
	var configuration models.Configuration
	err := r.db.First(&configuration, "name = ?", name).Error
	if err != nil {
		return models.Configuration{}, err
	}
	return configuration, nil
}

func (r *ConfigurationRepository) GetConfigurationTypeById(name string) (models.ConfigurationType, error) {
	var configurationType models.ConfigurationType
	err := r.db.First(&configurationType, "name = ?", name).Error
	if err != nil {
		return models.ConfigurationType{}, err
	}
	return configurationType, nil
}

func (r *ConfigurationRepository) ListConfigurations() ([]models.Configuration, error) {
	var configurations []models.Configuration
	err := r.db.Find(&configurations).Error
	if err != nil {
		return []models.Configuration{}, err
	}
	return configurations, nil
}

func (r *ConfigurationRepository) ListConfigurationTypes() ([]models.ConfigurationType, error) {
	var configurationTypes []models.ConfigurationType
	err := r.db.Find(&configurationTypes).Error
	if err != nil {
		return []models.ConfigurationType{}, err
	}
	return configurationTypes, nil
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
