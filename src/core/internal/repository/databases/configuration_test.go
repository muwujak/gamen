package databases

import (
	"testing"

	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/services"
)

func TestConfigurationRepositoryTypeAssertion(t *testing.T) {
	var _ interfaces.IConfigurationService = &ConfigurationRepository{}
}
