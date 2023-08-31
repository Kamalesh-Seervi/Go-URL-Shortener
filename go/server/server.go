package server

import (
	"strconv"

	"github.com/Kamalesh-Seervi/url/models"
	"github.com/Kamalesh-Seervi/url/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getAllRedirects(c *gin.Context) {
	urls, err := models.GetAllUrl()
	if err != nil {
		c.JSON(503, gin.H{"message": "Internal Server Error"})
		return
	}
	c.JSON(200, urls)
}

func getUrl(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(503, gin.H{"message": "Internal Server Error"})
		return
	}

	url, err := models.GetOneUrl(id)
	if err != nil {
		c.JSON(503, gin.H{"message": "Could Not return Url from DB"})
		return
	}

	c.JSON(200, url)
}
func createUrl(c *gin.Context) {
	c.Request.ParseForm()
	url := models.Url{
		Url:    uuid.New().String(),
		Random: c.Request.FormValue("random") == "true",
	}

	if url.Random {
		url.Url = utils.Base62(8)
	}

	err := models.CreateUrl(url)
	if err != nil {
		c.JSON(500, gin.H{"message": "could not create goly in db"})
		return
	}
	c.JSON(201, url)
}

func ServerListen() {
	router := gin.Default()

	router.Use(cors.Default())
	router.GET("/urls", getAllRedirects)
	router.GET("/url/:id", getUrl)
	router.POST("/url", createUrl)

	router.Run()
}

