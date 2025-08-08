package databases

import (
	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindTeamMemberByUserIdAndTeamId(userId uuid.UUID, teamId uuid.UUID) (models.TeamMember, error) {
	var teamMember models.TeamMember
	if err := r.db.Where("user_id = ? AND team_id = ?", userId, teamId).First(&teamMember).Error; err != nil {
		return models.TeamMember{}, err
	}
	return teamMember, nil
}
