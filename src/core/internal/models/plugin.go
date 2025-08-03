package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type PluginType string

const (
	CatalogueTypeForm   PluginType = "form"
	CatalogueTypeGraph  PluginType = "graph"
	CatalogueTypeAction PluginType = "action"
)

func (c PluginType) String() string {
	return string(c)
}

func (c *PluginType) Scan(value interface{}) error {
	if value == nil {
		*c = ""
		return nil
	}
	if str, ok := value.(string); ok {
		*c = PluginType(str)
		return nil
	}
	return fmt.Errorf("cannot scan %T into CatalogueType", value)
}

func (c PluginType) Value() (driver.Value, error) {
	return string(c), nil
}

type PluginTag struct {
	ID          uuid.UUID `json:"id" db:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `json:"name" db:"name" gorm:"uniqueIndex;not null"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	PluginTagMappings []PluginTagMapping `json:"plugin_mappings,omitempty" gorm:"foreignKey:TagID"`
}

type Plugin struct {
	ID                  uuid.UUID  `json:"id" db:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name                string     `json:"name" db:"name" gorm:"uniqueIndex;not null"`
	Description         string     `json:"description" db:"description"`
	Type                PluginType `json:"type" db:"type" gorm:"type:varchar(20);not null"`
	UISchema            JSON       `json:"ui_schema" db:"ui_schema" gorm:"type:jsonb"`
	Version             string     `json:"version" db:"version"`
	IsActive            bool       `json:"is_active" db:"is_active" gorm:"default:true"`
	ConfigurationTypeID uuid.UUID  `json:"configuration_type_id" db:"configuration_type_id" gorm:"type:uuid;not null"`
	CreatedAt           time.Time  `json:"created_at" db:"created_at" gorm:"autoCreateTime"`
	UpdatedAt           time.Time  `json:"updated_at" db:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	ConfigurationType ConfigurationType  `json:"configuration_type,omitempty" gorm:"foreignKey:ConfigurationTypeID"`
	PluginTagMappings []PluginTagMapping `json:"tag_mappings,omitempty" gorm:"foreignKey:PluginID"`
	Widgets           []Widget           `json:"widgets,omitempty" gorm:"foreignKey:PluginID"`
}

type PluginTagMapping struct {
	ID        uuid.UUID `json:"id" db:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	PluginID  uuid.UUID `json:"plugin_id" db:"plugin_id" gorm:"type:uuid;not null"`
	TagID     uuid.UUID `json:"tag_id" db:"tag_id" gorm:"type:uuid;not null"`
	CreatedAt time.Time `json:"created_at" db:"created_at" gorm:"autoCreateTime"`

	// Relationships
	Plugin Plugin    `json:"plugin,omitempty" gorm:"foreignKey:PluginID"`
	Tag    PluginTag `json:"tag,omitempty" gorm:"foreignKey:TagID"`
}
