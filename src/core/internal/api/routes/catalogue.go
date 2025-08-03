package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mujak27/gamen/src/core/internal/api/handlers"
)

func RegisterCatalogueRoutes(router *gin.RouterGroup, catalogueHandler *handlers.CatalogueHandler) {
	catalogueGroup := router.Group("/catalogues")
	{
		catalogueGroup.GET("", catalogueHandler.ListCatalogueKeys)
		catalogueGroup.GET("/:id", catalogueHandler.GetCatalogueModelById)
		catalogueGroup.GET("/handler/:id")
	}
}
