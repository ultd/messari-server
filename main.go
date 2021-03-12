package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/ultd/messari-server/handlers"
)

func main() {
	apiKey := os.Getenv("MESSARI_API_KEY")

	if apiKey == "" {
		panic("need MESSARI_API_KEY in env in order to run server!")
	}

	// setting debug level logging
	logrus.SetLevel(logrus.DebugLevel)

	server := gin.Default()

	server.GET("/api/asset", handlers.GetAllAssetsHandler(apiKey))
	server.GET("/api/asset/:symbolOrSlug", handlers.GetAssetMetricsHandler(apiKey))
	server.GET("/api/aggregate", handlers.GetAssetMetricsAggregateHandler(apiKey))

	err := server.Run(":8000")
	if err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
