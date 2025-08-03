package databases

import (
	"github.com/mujak27/gamen/src/core/internal/models"
	"gorm.io/gorm"
)

type WidgetRepository struct {
	db *gorm.DB
}

func NewWidgetRepository(db *gorm.DB) *WidgetRepository {
	return &WidgetRepository{db: db}
}

func (r *WidgetRepository) GetWidgetById(id string) (models.Widget, error) {
	var widget models.Widget
	err := r.db.First(&widget, "id = ?", id).Error
	if err != nil {
		return models.Widget{}, err
	}
	return widget, nil
}
