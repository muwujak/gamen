package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/services"
)

type CatalogueHandler struct {
	service interfaces.CatalogueService
}

func NewCatalogueHandler(
	service interfaces.CatalogueService,
) *CatalogueHandler {
	return &CatalogueHandler{
		service: service,
	}
}

func (catalogueHandler *CatalogueHandler) GetCatalogueModelById(ctx *gin.Context) {
	id := ctx.Param("id")
	catalogue, err := catalogueHandler.service.GetPluginModelById(uuid.MustParse(id))
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve catalogue"})
		return
	}
	ctx.JSON(200, gin.H{"catalogue": catalogue})
}

func (catalogueHandler *CatalogueHandler) ListCatalogueKeys(ctx *gin.Context) {
	catalogueType, err := catalogueHandler.service.ListPluginKeys()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve catalogue"})
		return
	}
	ctx.JSON(200, gin.H{"catalogue": catalogueType})
}
