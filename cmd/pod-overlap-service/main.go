package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	httpapi "github.com/brunooliveiramac/pod-overlap-service/internal/platform/http"
)

func main() {
	router := gin.Default()
	httpapi.RegisterRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("overlap-service is running on :%s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
