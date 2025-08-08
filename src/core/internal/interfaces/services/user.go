package interfaces

import (
	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/api/dto"
	"github.com/mujak27/gamen/src/core/internal/models"
)

type IUserService interface {
	Login(dto.LoginRequest) (*dto.LoginResponse, error)
	Register(dto.RegisterRequest) (*models.User, error)
	IsTeamMember(userId uuid.UUID, teamId uuid.UUID) (bool, error)
}
