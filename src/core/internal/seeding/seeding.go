package seeding

import (
	"github.com/mujak27/gamen/src/core/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Configuration{})
}
