package services

import (
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/repository"
	"github.com/mujak27/gamen/src/core/internal/models"
)

type DashboardService struct {
	repo interfaces.DashboardRepository
}

func NewDashboardService(repo interfaces.DashboardRepository) *DashboardService {
	return &DashboardService{repo: repo}
}

func (s *DashboardService) GetCurrentUserDashboard() (models.Dashboard, error) {
	data, err := s.repo.GetCurrentUserDashboard()
	if err != nil {
		return models.Dashboard{}, err
	}
	return data, nil
}
