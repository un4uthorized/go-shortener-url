package handlers

import (
	"go-shortener-url/internal/app/entities"
	"go-shortener-url/internal/app/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShortenHandler struct {
	ShortenURLUseCase *usecases.ShortenURLUseCase
}

func (h *ShortenHandler) ShortenURL(c *gin.Context) {
	var request entities.ShortenURLRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortenedURL, err := h.ShortenURLUseCase.Execute(request.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"shortened_url": shortenedURL})
}
