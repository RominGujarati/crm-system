package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"crm-system/api/models"
	"crm-system/api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var interactionCollection *mongo.Collection

func init() {
	interactionCollection = utils.GetMongoClient().Database("crm").Collection("interactions")
}

func GetInteractions(c *gin.Context) {
	cursor, err := interactionCollection.Find(context.Background(), bson.M{})
	if err != nil {
		utils.LogError("Failed to fetch interactions", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch interactions"})
		return
	}
	defer cursor.Close(context.Background())

	var interactions []models.Interaction
	if err := cursor.All(context.Background(), &interactions); err != nil {
		utils.LogError("Failed to decode interactions", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode interactions"})
		return
	}

	c.JSON(http.StatusOK, interactions)
}

func CreateInteraction(c *gin.Context) {
	var interaction models.Interaction
	if err := c.ShouldBindJSON(&interaction); err != nil {
		utils.LogError("Invalid interaction data", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interaction data"})
		return
	}

	_, err := interactionCollection.InsertOne(context.Background(), interaction)
	if err != nil {
		utils.LogError("Failed to create interaction", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create interaction"})
		return
	}

	c.JSON(http.StatusCreated, interaction)
}

func UpdateInteraction(c *gin.Context) {
	id := c.Param("id")
	var interaction models.Interaction
	if err := c.ShouldBindJSON(&interaction); err != nil {
		utils.LogError("Invalid interaction data", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interaction data"})
		return
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": interaction}
	_, err := interactionCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		utils.LogError("Failed to update interaction", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update interaction"})
		return
	}

	c.JSON(http.StatusOK, interaction)
}

func DeleteInteraction(c *gin.Context) {
	id := c.Param("id")
	_, err := interactionCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		utils.LogError("Failed to delete interaction", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete interaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Interaction deleted"})
}
