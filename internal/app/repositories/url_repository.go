package repositories

import "go-shortener-url/internal/app/entities"

type URLRepository interface {
	SaveURL(url *entities.URL) error
	GetURLByID(ID string) (*entities.URL, error)
}
