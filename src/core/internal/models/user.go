package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Enums
type TeamMemberRole string

const (
	TeamMemberRoleAdmin  TeamMemberRole = "admin"
	TeamMemberRoleMember TeamMemberRole = "member"
	TeamMemberRoleViewer TeamMemberRole = "viewer"
)

func (r TeamMemberRole) String() string {
	return string(r)
}

func (r *TeamMemberRole) Scan(value interface{}) error {
	if value == nil {
		*r = ""
		return nil
	}
	if str, ok := value.(string); ok {
		*r = TeamMemberRole(str)
		return nil
	}
	return fmt.Errorf("cannot scan %T into TeamMemberRole", value)
}

func (r TeamMemberRole) Value() (driver.Value, error) {
	return string(r), nil
}

type User struct {
	ID           uuid.UUID  `json:"id" db:"id" gorm:"type:uuid;primary_key;not null"`
	Email        string     `json:"email" db:"email" gorm:"uniqueIndex;not null"`
	Username     string     `json:"username" db:"username" gorm:"uniqueIndex;not null"`
	FirstName    string     `json:"first_name" db:"first_name"`
	LastName     string     `json:"last_name" db:"last_name"`
	PasswordHash string     `json:"-" db:"password_hash" gorm:"not null"`
	IsActive     bool       `json:"is_active" db:"is_active" gorm:"default:true"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at" gorm:"autoUpdateTime"`
	LastLogin    *time.Time `json:"last_login" db:"last_login"`

	// Relationships
	CreatedTeams      []Team          `json:"created_teams,omitempty" gorm:"foreignKey:CreatedBy"`
	TeamMemberships   []TeamMember    `json:"team_memberships,omitempty" gorm:"foreignKey:UserID"`
	CreatedConfigs    []Configuration `json:"created_configs,omitempty" gorm:"foreignKey:CreatedBy"`
	CreatedDashboards []Dashboard     `json:"created_dashboards,omitempty" gorm:"foreignKey:CreatedBy"`
	CreatedWidgets    []Widget        `json:"created_widgets,omitempty" gorm:"foreignKey:CreatedBy"`
}

type Team struct {
	ID          uuid.UUID `json:"id" db:"id" gorm:"type:uuid;primary_key;not null"`
	Name        string    `json:"name" db:"name" gorm:"uniqueIndex;not null"`
	Description string    `json:"description" db:"description"`
	IsActive    bool      `json:"is_active" db:"is_active" gorm:"default:true"`
	CreatedBy   uuid.UUID `json:"created_by" db:"created_by" gorm:"type:uuid;not null"`
	CreatedAt   time.Time `json:"created_at" db:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	Creator        User            `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
	Members        []TeamMember    `json:"members,omitempty" gorm:"foreignKey:TeamID"`
	Configurations []Configuration `json:"configurations,omitempty" gorm:"foreignKey:TeamID"`
	Dashboards     []Dashboard     `json:"dashboards,omitempty" gorm:"foreignKey:TeamID"`
}

type TeamMember struct {
	ID        uuid.UUID      `json:"id" db:"id" gorm:"type:uuid;primary_key;not null"`
	TeamID    uuid.UUID      `json:"team_id" db:"team_id" gorm:"type:uuid;not null"`
	UserID    uuid.UUID      `json:"user_id" db:"user_id" gorm:"type:uuid;not null"`
	Role      TeamMemberRole `json:"role" db:"role" gorm:"type:varchar(20);not null"`
	IsActive  bool           `json:"is_active" db:"is_active" gorm:"default:true"`
	JoinedAt  time.Time      `json:"joined_at" db:"joined_at" gorm:"autoCreateTime"`
	CreatedAt time.Time      `json:"created_at" db:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at" gorm:"autoUpdateTime"`

	// Relationships
	Team Team `json:"team,omitempty" gorm:"foreignKey:TeamID"`
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
