package databases

import (
	"testing"

	"github.com/google/uuid"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/repository"
	"github.com/mujak27/gamen/src/core/internal/seeding"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestConfigurationRepositoryTypeAssertion(t *testing.T) {
	var _ interfaces.IConfigurationRepository = &ConfigurationRepository{}
}

func TestConfigurationRepositoryGetConfigurationById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("mock.db"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	seeding.Migrate(db)
	configurationRepo := NewConfigurationRepository(db)
	configurationRepo.GetConfigurationById(uuid.New())
}
