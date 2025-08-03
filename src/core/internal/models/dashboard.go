package models

import (
	"time"

	"github.com/google/uuid"
)

type Dashboard struct {
	ID             uuid.UUID  `json:"id" db:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name           string     `json:"name" db:"name" gorm:"not null"`
	Description    string     `json:"description" db:"description"`
	TeamID         uuid.UUID  `json:"team_id" db:"team_id" gorm:"type:uuid;not null"`
	LayoutConfig   JSON       `json:"layout_config" db:"layout_config" gorm:"type:jsonb"`
	IsDefault      bool       `json:"is_default" db:"is_default" gorm:"default:false"`
	CreatedBy      uuid.UUID  `json:"created_by" db:"created_by" gorm:"type:uuid;not null"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at" gorm:"autoUpdateTime"`
	LastAccessedAt *time.Time `json:"last_accessed_at" db:"last_accessed_at"`

	// Relationships
	Team    Team     `json:"team,omitempty" gorm:"foreignKey:TeamID"`
	Creator User     `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
	Widgets []Widget `json:"widgets,omitempty" gorm:"foreignKey:DashboardID"`
}

type Widget struct {
	ID              uuid.UUID  `json:"id" db:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	DashboardID     uuid.UUID  `json:"dashboard_id" db:"dashboard_id" gorm:"type:uuid;not null"`
	PluginID        uuid.UUID  `json:"plugin_id" db:"plugin_id" gorm:"type:uuid;not null"`
	ConfigurationID uuid.UUID  `json:"configuration_id" db:"configuration_id" gorm:"type:uuid;not null"`
	Name            string     `json:"name" db:"name" gorm:"not null"`
	Description     string     `json:"description" db:"description"`
	IsActive        bool       `json:"is_active" db:"is_active" gorm:"default:true"`
	CreatedBy       uuid.UUID  `json:"created_by" db:"created_by" gorm:"type:uuid;not null"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time  `json:"updated_at" db:"updated_at" gorm:"autoUpdateTime"`
	LastExecutedAt  *time.Time `json:"last_executed_at" db:"last_executed_at"`

	// Relationships
	Dashboard     Dashboard     `json:"dashboard,omitempty" gorm:"foreignKey:DashboardID"`
	Plugin        Plugin        `json:"plugin,omitempty" gorm:"foreignKey:PluginID"`
	Configuration Configuration `json:"configuration,omitempty" gorm:"foreignKey:ConfigurationID"`
	Creator       User          `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
}
