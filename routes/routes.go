package routes


import (
	"github.com/gin-gonic/gin"

	"url-shortener/controllers"
)


func SetupRoutes(router *gin.Engine){


	router.POST(
		"/shorten",
		controllers.CreateShortURL,
	)

}