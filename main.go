package main

import (
	"fmt"
	"log"
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

	db, err := mongod.MongoConnection(secrets.DatabaseURL, secrets.DatabaseName)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	logger := logrus.New()
	
	services := services.NewService(db, logger)
	controller := controllers.NewController(services)
	handlers := routes.NewRoute(controller)


	routes.Routes(router, *handlers)

	log.Fatal(router.Run(":" + port))
}
