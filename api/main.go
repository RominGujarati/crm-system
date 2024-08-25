package main

import (
	"github.com/gin-gonic/gin"
	"crm-system/api/routes"
	"crm-system/api/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
	// Initialize MongoDB client
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(nil, clientOptions)
	if err != nil {
		utils.LogError("Failed to connect to MongoDB", err)
		return
	}

	r := gin.Default()
	r.Use(utils.AuthMiddleware())
	routes.SetupUserRoutes(r)
	routes.SetupCustomerRoutes(r)
	routes.SetupAuthRoutes(r)
	routes.SetupInteractionRoutes(r)
	r.Run()
}

func GetMongoClient() *mongo.Client {
	return client
}
