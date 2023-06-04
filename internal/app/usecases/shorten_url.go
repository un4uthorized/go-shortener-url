package usecases

import (
	"encoding/base64"
	"go-shortener-url/internal/app/entities"
	"go-shortener-url/internal/app/repositories"
	"go-shortener-url/pkg/utils"
)

type ShortenURLUseCase struct {
	Repository  repositories.URLRepository
	IDGenerator utils.IDGenerator
}

func (uc *ShortenURLUseCase) Execute(originalURL string) (string, error) {
	id := uc.IDGenerator.GenerateID()

	shortened := base64.StdEncoding.EncodeToString([]byte(id))

	url := &entities.URL{
		ID:          id,
		OriginalURL: originalURL,
	}

	if err := uc.Repository.SaveURL(url); err != nil {
		return "", err
	}

	return shortened, nil
}
