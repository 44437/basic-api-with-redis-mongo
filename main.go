package main

import (
	"context"
	"log"
	"os"
	"redis/cache"
	"redis/controller"
	"redis/mongodb"
	"redis/repository"
	"redis/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func init() {
	err := godotenv.Load(".config/.env")
	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	const (
		cacheDB     = 1
		cacheExpire = 1
	)
	var err error

	mongoDB, err := mongodb.NewMongoDB()
	if err != nil {
		log.Panicln(err)
	}
	defer mongoDB.CloseConnection(context.Background())

	serverRedisCache := cache.NewRedisCache(
		os.Getenv("CACHE_HOST"),
		os.Getenv("CACHE_PASSWORD"),
		cacheDB,
		cacheExpire,
	)
	serverRepository := repository.NewRepository(mongoDB)
	serverService := service.NewService(serverRepository, serverRedisCache)
	serverController := controller.NewController(serverService)

	server := echo.New()
	server.HideBanner = true
	server.GET("/humans", serverController.GetHumans)
	server.GET("/humans/:id", serverController.GetHuman)
	err = server.Start(":8080")
	if err != nil {
		log.Panicln(err)
	}
}
