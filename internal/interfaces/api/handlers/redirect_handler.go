package handlers

import (
	"go-shortener-url/internal/app/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OriginalUrlHandler struct {
	OriginalURLUseCase *usecases.OriginalURLUseCase
}

func (h *OriginalUrlHandler) RedirectToOriginalURL(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	originalURL, err := h.OriginalURLUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}
