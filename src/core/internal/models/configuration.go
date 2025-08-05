package models

import (
	"time"

	"github.com/google/uuid"
)

type ConfigurationType struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Schema      JSON      `json:"schema"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Configuration struct {
	ID                  uuid.UUID  `json:"id" db:"id" gorm:"type:uuid;primary_key;not null"`
	Name                string     `json:"name" db:"name" gorm:"not null"`
	Description         string     `json:"description" db:"description"`
	ConfigurationTypeID uuid.UUID  `json:"configuration_type_id" db:"configuration_type_id" gorm:"type:uuid;not null"`
	TeamID              uuid.UUID  `json:"team_id" db:"team_id" gorm:"type:uuid;not null"`
	IsActive            bool       `json:"is_active" db:"is_active" gorm:"default:true"`
	CreatedBy           uuid.UUID  `json:"created_by" db:"created_by" gorm:"type:uuid;not null"`
	CreatedAt           time.Time  `json:"created_at" db:"created_at" gorm:"autoCreateTime"`
	UpdatedAt           time.Time  `json:"updated_at" db:"updated_at" gorm:"autoUpdateTime"`
	LastUsedAt          *time.Time `json:"last_used_at" db:"last_used_at"`
	Data                JSON       `json:"data" db:"data" gorm:"type:jsonb"`

	// Relationships
	Team    Team     `json:"team,omitempty" gorm:"foreignKey:TeamID"`
	Creator User     `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
	Widgets []Widget `json:"widgets,omitempty" gorm:"foreignKey:ConfigurationID"`
}
