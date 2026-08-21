package main

import (
	"polaris/internal/db"
	"polaris/internal/embedding"
	"polaris/internal/gateway/controllers"
	"polaris/internal/pipeline"

	"github.com/gin-gonic/gin"
)

func main() {
	// database
	dbService, err := db.New("TODO")
	if err != nil {
		panic(err)
	}

	// embedding
	embeddingService := embedding.New("TODO", 0)

	// pipeline
	mainService := pipeline.New(embeddingService, dbService)

	// router
	router := gin.Default()
	routingService := controllers.NewService(mainService)

	fileRoutes := router.Group("/content")
	fileRoutes.POST("/ingest", routingService.Ingest)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
