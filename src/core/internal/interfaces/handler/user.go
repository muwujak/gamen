package interfaces

import "github.com/gin-gonic/gin"

type UserHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}
