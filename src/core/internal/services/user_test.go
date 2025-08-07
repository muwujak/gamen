package services

import (
	"testing"

	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/services"
)

func TestUserServiceTypeAssertion(t *testing.T) {
	var _ interfaces.UserService = &UserService{}
}
