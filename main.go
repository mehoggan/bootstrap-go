package main

import (
	"github.com/gin-gonic/gin"

	"github.com/mehoggan/vinyl-collection-service-go/endpoints"
)

func main() {
	router := gin.Default()
	router.GET("albums", endpoints.GetAlbumsHandler)
	router.POST("albums", endpoints.PostAlbumsHandler)

	router.Run("localhost:8080")
}
