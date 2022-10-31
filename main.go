package main

import (
	"github.com/gin-gonic/gin"

	"github.com/mehoggan/vinyl-collection-service-go/endpoints"
)

func main() {
	router := gin.Default()
	router.GET("albums", endpoints.GetAlbumsHandler)

	router.Run("localhost:8080")
}
