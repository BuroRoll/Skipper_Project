package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetCatalog(c *gin.Context) {
	catalog := h.services.GetCatalog()
	c.JSON(http.StatusOK, gin.H{
		"catalog": catalog,
	})
	//c.JSON(http.StatusOK, catalog)
}

func (h *Handler) GetMainSection(c *gin.Context) {
	mainCatalog := h.services.GetMainCatalog()
	c.JSON(http.StatusOK, gin.H{
		"Main_catalog": mainCatalog,
	})
}

func (h *Handler) GetCatalogChild(c *gin.Context) {
	catalogChild, err := h.services.GetCatalogChild()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка полчения данных",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"catalog_child": catalogChild,
	})
}

func (h *Handler) GetClasses(c *gin.Context) {
	//userId := h.getAuthStatus(c)
	pagination := GeneratePaginationFromRequest(c)

	//if userId != 0 {
	//	isFavouriteUser = h.services.IsFavouriteUser(userId, mentorId)
	//} else {
	userLists, err := h.services.GetClasses(&pagination)
	//}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"catalog_of_mentors": userLists,
	})
}
