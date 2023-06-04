package persistence

import (
	"go-shortener-url/internal/app/entities"

	"github.com/go-redis/redis"
)

type RedisDB struct {
	client *redis.Client
}

func NewRedisDB() *RedisDB {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &RedisDB{
		client: client,
	}
}

func (r *RedisDB) SaveURL(url *entities.URL) error {
	err := r.client.Set(url.ID, url.OriginalURL, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisDB) GetURLByID(ID string) (*entities.URL, error) {
	result, err := r.client.Get(ID).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	url := &entities.URL{
		ID:          ID,
		OriginalURL: result,
	}

	return url, nil
}
