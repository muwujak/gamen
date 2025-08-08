package services

import (
	"testing"

	"github.com/google/uuid"
	repositoryMocks "github.com/mujak27/gamen/src/core/internal/interfaces/repository/mocks"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/services"
	"github.com/mujak27/gamen/src/core/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceTypeAssertion(t *testing.T) {
	var _ interfaces.IUserService = &UserService{}
}

func TestUserServiceIsTeamMember(t *testing.T) {
	userIdA := uuid.New()
	userIdB := uuid.New()
	teamId := uuid.New()

	mockUserRepository := repositoryMocks.NewMockUserRepository(t)
	userService := NewUserService(mockUserRepository)

	// CASE: user is a team member
	mockUserRepository.EXPECT().FindTeamMemberByUserIdAndTeamId(userIdA, teamId).Return(models.TeamMember{
		ID: uuid.New(),
	}, nil).Once()
	isTeamMember, err := userService.IsTeamMember(userIdA, teamId)
	assert.NoError(t, err)
	assert.True(t, isTeamMember)

	// CASE: user is not a team member
	mockUserRepository.EXPECT().FindTeamMemberByUserIdAndTeamId(userIdB, teamId).Return(models.TeamMember{}, nil).Once()
	isTeamMember, err = userService.IsTeamMember(userIdB, teamId)
	assert.NoError(t, err)
	assert.False(t, isTeamMember)
}
