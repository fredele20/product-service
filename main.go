package main

import (
	"fmt"
	"log"
	"product-service/setup"
	"product-service/source/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// secrets := config.GetSecrets()
	// port := secrets.Port
	address := fmt.Sprintf("127.0.0.1:%s", setup.Secrets.Port)

	// db, err := mongod.MongoConnection(secrets.DatabaseURL, secrets.DatabaseName)
	// if err != nil {
	// 	fmt.Println(err)
	// 	log.Fatal(err)
	// }

	// redis, err := cache.RedisConnection(secrets.RedisUrl)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	router := gin.New()
	router.Use(gin.Logger())
	// logger := logrus.New()

	routes.RouteHandlers(router)

	// services := services.NewService(db, logger, redis)
	// controller := controllers.NewController(services)
	// handlers := routes.NewRoute(controller)

	// routes.Routes(router, *handlers)

	log.Fatal(router.Run(address))
}
