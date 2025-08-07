package routes

import (
	"github.com/gin-gonic/gin"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/handler"
)

func RegisterUserRoutes(router *gin.RouterGroup, userHandler interfaces.UserHandler) {

	router.POST("/login", userHandler.Login)
	router.POST("/register", userHandler.Register)
}
