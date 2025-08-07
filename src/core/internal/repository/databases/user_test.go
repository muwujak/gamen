package databases

import (
	"testing"

	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/repository"
)

func TestUserRepositoryTypeAssertion(t *testing.T) {
	var _ interfaces.UserRepository = &UserRepository{}
}
