package databases

import "github.com/mujak27/gamen/src/core/internal/models"

type DashboardRepository struct {
	db interface{}
}

func NewDashboardRepository(db interface{}) *DashboardRepository {
	return &DashboardRepository{db: db}
}

func (r *DashboardRepository) GetCurrentUserDashboard() (models.Dashboard, error) {
	return models.Dashboard{}, nil
}
