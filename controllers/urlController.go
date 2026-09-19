package controllers


import (
	"net/http"

	"github.com/gin-gonic/gin"

	"url-shortener/config"
	"url-shortener/models"
	"url-shortener/utils"

	"time"
)



func CreateShortURL(c *gin.Context){


	var body struct{

		URL string `json:"url"`

	}


	if err:=c.ShouldBindJSON(&body); err!=nil{

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error":"Invalid request",
			},
		)

		return

	}



	code:=utils.GenerateCode(6)



	url:=models.URL{

		OriginalURL:body.URL,

		ShortCode:code,

		Clicks:0,

		CreatedAt:time.Now(),

	}



	_,err:=config.URLCollection.InsertOne(
		c,
		url,
	)



	if err!=nil{

		c.JSON(500,gin.H{
			"error":"Database error",
		})

		return
	}



	c.JSON(201,gin.H{

		"short_url":
		"http://localhost:5000/"+code,

		"code":code,

	})


}