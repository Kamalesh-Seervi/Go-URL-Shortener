package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Kamalesh-Seervi/url/models"
	"github.com/Kamalesh-Seervi/url/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getAllUrl(c *gin.Context) {
	urls, err := models.GetAllUrl()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve URLs"})
		return
	}
	c.JSON(200, urls)
}

func redirectLink(c *gin.Context) {
	golyURL := c.Param("redirect")
	goly, err := models.FindByGolyUrl(golyURL)
	fmt.Println("golyURL:", golyURL)
	fmt.Println("URl", goly)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not find goly in DB " + err.Error(),
		})
		return
	}

	// Grab any stats you want...
	goly.Clicked += 1

	err = models.UpdateUrl(goly)
	if err != nil {
		fmt.Printf("error updating: %v\n", err)
	}

	c.Redirect(http.StatusTemporaryRedirect, goly.Redirect)
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
	var url models.Url

	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error parsing JSON " + err.Error(),
		})
		return
	}

	if url.Random {
		url.Url = utils.Base62(8)
	}

	if err := models.CreateUrl(&url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not create goly in db " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, url)
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

	router.GET("/:redirect", redirectLink)
	router.GET("/urls", getAllUrl)
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
