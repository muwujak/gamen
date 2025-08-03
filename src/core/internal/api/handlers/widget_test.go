package handlers

import (
	"testing"

	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/handler"
)

func TestWidgetHandlerTypeAssertion(t *testing.T) {
	var _ interfaces.WidgetHandler = &WidgetHandler{}
}

// mockWidgetService := func() interfaces.WidgetService {
// 	return &services.WidgetService{
// 		WidgetRepository: &repositories.WidgetRepository{},
// 	}
// }

func TestWidgetHandlerAction(t *testing.T) {

}
