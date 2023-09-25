package main

import (
	"fmt"
	"log"
	"product-service/cache"
	"product-service/config"
	"product-service/controllers"
	"product-service/database/mongod"
	"product-service/routes"
	"product-service/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	secrets := config.GetSecrets()
	port := secrets.Port
	address := fmt.Sprintf("127.0.0.1:%s", port)

	db, err := mongod.MongoConnection(secrets.DatabaseURL, secrets.DatabaseName)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	redis, err := cache.RedisConnection(secrets.RedisUrl)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	logger := logrus.New()

	services := services.NewService(db, logger, redis)
	controller := controllers.NewController(services)
	handlers := routes.NewRoute(controller)

	routes.Routes(router, *handlers)

	log.Fatal(router.Run(address))
}
