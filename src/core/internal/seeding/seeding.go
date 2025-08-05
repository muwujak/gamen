package seeding

import (
	"github.com/mujak27/gamen/src/core/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Configuration{})
	db.AutoMigrate(&models.PluginTag{})
	db.AutoMigrate(&models.PluginTagMapping{})
	db.AutoMigrate(&models.Widget{})
	db.AutoMigrate(&models.Dashboard{})
	db.AutoMigrate(&models.Team{})
	db.AutoMigrate(&models.TeamMember{})
}
