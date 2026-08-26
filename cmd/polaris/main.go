package main

import (
	"polaris/internal/config"
	"polaris/internal/db"
	"polaris/internal/embedding"
	"polaris/internal/gateway/controllers"
	"polaris/internal/pipeline"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// config
	cfg, err := config.LoadConfig("./.env")
	if err != nil {
		panic(err)
	}

	// database
	dbService, err := db.New(cfg)
	if err != nil {
		panic(err)
	}
	defer dbService.Close()

	// embedding
	embeddingService := embedding.New(cfg.EmbeddingURL, 10*time.Second)

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
