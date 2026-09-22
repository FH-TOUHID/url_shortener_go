package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"url-shortener/utils"
)
var urls = map[string]string{}

func shortenURL(c *gin.Context) {
	var body struct {
		URL string `json:"url"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	code := utils.GenerateCode(6)
	urls[code] = body.URL

	c.JSON(http.StatusCreated, gin.H{
		"short_url": "http://localhost:5000/" + code,
		"code":      code,
	})
}

func getURL(c *gin.Context) {
	code := c.Param("code")

	originalURL, found := urls[code]

	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "URL not found",
		})
		return
	}

	c.Redirect(http.StatusFound, originalURL)
}

func main() {
	router := gin.Default()

	router.POST("/shorten", shortenURL)
	router.GET("/:code", getURL)

	router.Run(":5000")
}