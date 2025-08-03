package interfaces

import "github.com/gin-gonic/gin"

type WidgetHandler interface {
	Action(ctx *gin.Context)
	// Fetch(ctx *gin.Context)
	// Get(ctx *gin.Context)
}
