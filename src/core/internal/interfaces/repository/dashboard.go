package interfaces

import (
	"github.com/mujak27/gamen/src/core/internal/models"
)

// REFACTOR
type DashboardRepository interface {
	GetCurrentUserDashboard() (models.Dashboard, error)
}
