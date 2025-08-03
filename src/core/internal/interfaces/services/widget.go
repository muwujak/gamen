package interfaces

import (
	"github.com/mujak27/gamen/src/core/internal/api/dto"
)

type WidgetService interface {
	Action(dto.WidgetActionPayload) error
}
