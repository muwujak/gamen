package interfaces

import (
	"github.com/mujak27/gamen/src/core/internal/api/dto"
	"github.com/mujak27/gamen/src/core/internal/models"
)

type UserService interface {
	Login(dto.LoginRequest) (*dto.LoginResponse, error)
	Register(dto.RegisterRequest) (*models.User, error)
}
