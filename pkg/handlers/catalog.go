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

func (h *Handler) getCatalog(c *gin.Context) {
	catalog := h.services.GetCatalog()
	c.JSON(http.StatusOK, gin.H{
		"catalog": catalog,
	})
}

func (h *Handler) mainSection(c *gin.Context) {
	mainCatalog := h.services.GetMainCatalog()
	c.JSON(http.StatusOK, gin.H{
		"Main_catalog": mainCatalog,
	})
}
