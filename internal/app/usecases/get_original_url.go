package usecases

import (
	"encoding/base64"
	"go-shortener-url/internal/app/repositories"
)

type OriginalURLUseCase struct {
	Repository repositories.URLRepository
}

func (uc *OriginalURLUseCase) Execute(id string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		return "", err
	}

	url, err := uc.Repository.GetURLByID(string(decoded))
	if err != nil {
		return "", err
	}

	return url.OriginalURL, nil
}
