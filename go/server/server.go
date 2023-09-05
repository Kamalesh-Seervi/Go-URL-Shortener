package server

import (
	"net/http"
	"os"
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
		c.JSON(500, gin.H{"error": "Failed to retrieve URLs"})
		return
	}
	c.JSON(200, urls)
}

func getUrl(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	url, err := models.GetOneUrl(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve URL by ID"})
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
		c.JSON(500, gin.H{"error": "Failed to create URL in DB"})
		return
	}
	c.JSON(201, url)
}

func updateUrl(c *gin.Context) {
	var url models.Url

	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not parse JSON: " + err.Error(),
		})
		return
	}

	if err := models.UpdateUrl(url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update goly link in DB: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, url)
}

func deleteUrl(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not parse ID from URL: " + err.Error(),
		})
		return
	}

	if err := models.DeleteUrl(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete from DB: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Goly deleted.",
	})
}

func ServerListen() {
	router := gin.Default()

	router.Use(cors.Default())
	router.GET("/urls", getAllRedirects)
	router.GET("/url/:id", getUrl)
	router.POST("/url", createUrl)
	router.PATCH("/url", updateUrl)
	router.DELETE("/url/:id", deleteUrl)

	port := os.Getenv("PORT")
	if port == "" {
		port = "7777" // Default port
	}
	router.Run(":" + port)
}
