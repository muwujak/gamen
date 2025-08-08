package interfaces

import (
	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/models"
)

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) error
	FindTeamMemberByUserIdAndTeamId(userId uuid.UUID, teamId uuid.UUID) (models.TeamMember, error)
}
