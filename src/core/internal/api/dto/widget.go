package dto

import "github.com/google/uuid"

type WidgetActionPayload struct {
	WidgetID uuid.UUID              `json:"widget_id"`
	Data     map[string]interface{} `json:"data"`
}
