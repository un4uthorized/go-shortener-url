package routers

import (
	"go-shortener-url/internal/app/repositories"
	"go-shortener-url/internal/app/usecases"
	"go-shortener-url/internal/interfaces/api/handlers"
	"go-shortener-url/pkg/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter(repository repositories.URLRepository) *gin.Engine {
	router := gin.Default()

	shortenHandler := &handlers.ShortenHandler{
		ShortenURLUseCase: &usecases.ShortenURLUseCase{
			Repository:  repository,
			IDGenerator: utils.IDGenerator{},
		},
	}

	originalURLHandler := &handlers.OriginalUrlHandler{
		OriginalURLUseCase: &usecases.OriginalURLUseCase{
			Repository: repository,
		},
	}

	router.POST("/shorten", shortenHandler.ShortenURL)
	router.GET("/:id", originalURLHandler.RedirectToOriginalURL)

	return router
}
