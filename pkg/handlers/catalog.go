package handlers

import (
	"Skipper/pkg/models/forms"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createCatalog(c *gin.Context) {
	var catalog forms.CatalogInput
	if err := c.BindJSON(&catalog); err != nil {
		c.JSON(http.StatusBadRequest, "invalid input body")
		return
	}
	catalogId, err := h.services.CreateCatalog(catalog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"catalog_id": catalogId,
	})
}

func (h *Handler) getAllCatalog(c *gin.Context) {
	data := h.services.GetAllCatalog(1)
	c.JSON(http.StatusOK, gin.H{
		"catalog": data,
	})
}
