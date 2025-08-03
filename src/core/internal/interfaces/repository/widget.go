package interfaces

import (
	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/models"
)

// REFACTOR
type WidgetRepository interface {
	GetWidgetById(id uuid.UUID) (models.Widget, error)
}
