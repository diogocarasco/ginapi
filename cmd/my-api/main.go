package main

import (
	"github.com/diogocarasco/ginapi/api/routes"
	"github.com/diogocarasco/ginapi/internal/repository"
	"github.com/diogocarasco/ginapi/pkg/logger"
)

func main() {

	logger.Init()

	logger.Info("Starting API Server")

	logger.Info("Connecting to MongoDB")
	mongoRepo, err := repository.NewMongoRepository("mongodb+srv://carasco:carasco@clustergoapi.92n6j.mongodb.net/?retryWrites=true&w=majority&appName=ClusterGoApi", "sample_mflix")
	if err != nil {
		logger.Error("Couldn't connect to the MongoDB repository", err)
	}

	router := routes.NewRouter()
	if err := router.Run(":8080"); err != nil {
		logger.Fatal("FAILED: " + err.Error())
	}

}
