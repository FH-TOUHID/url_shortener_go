package main


import (

	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

	"url-shortener/config"

	"url-shortener/routes"

)



func main(){


	// load .env

	err := godotenv.Load()

	if err != nil {

		panic("env file not loaded")

	}



	// connect mongodb

	config.ConnectMongo()



	// create gin server

	router := gin.Default()



	// register routes

	routes.SetupRoutes(router)



	port := os.Getenv(
		"SERVER_PORT",
	)



	router.Run(
		":"+port,
	)


}