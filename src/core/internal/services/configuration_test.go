package services

import (
	"testing"

	serviceInterface "github.com/mujak27/gamen/src/core/internal/interfaces/services"
)

func TestConfigurationServiceTypeAssertion(t *testing.T) {
	var _ serviceInterface.ConfigurationService = &ConfigurationService{}
}
