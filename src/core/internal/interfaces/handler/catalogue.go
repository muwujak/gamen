package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PluginHandler interface {
	GetId() uuid.UUID
	Action(ctx *gin.Context)
	Fetch(ctx *gin.Context)
	Get(ctx *gin.Context)
}
