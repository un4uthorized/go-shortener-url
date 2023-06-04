package main

import (
	"go-shortener-url/internal/interfaces/api/routers"
	"go-shortener-url/internal/interfaces/persistence"
)

func main() {
	redisClient := persistence.NewRedisDB()
	router := routers.SetupRouter(redisClient)
	router.Run(":8080")
}
