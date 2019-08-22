package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	rm "github.com/khmerlang/Khmer2Romance/models/romanization"
)

type KhmerWords struct {
	Sentences string `json:"sentences"`
}

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		api.POST("/kh2rm", Kh2RmHandler)
	}

	// Start and run the server
	router.Run(":3000")
}

// LikeJoke increments the likes of a particular joke Item
func Kh2RmHandler(c *gin.Context) {
	var sentences KhmerWords
	c.BindJSON(&sentences)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": rm.Romanize(sentences.Sentences),
	})
}
