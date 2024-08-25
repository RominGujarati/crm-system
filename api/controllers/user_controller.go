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

var userCollection *mongo.Collection

func init() {
	userCollection = utils.GetMongoClient().Database("crm").Collection("users")
}

func GetUsers(c *gin.Context) {
	cursor, err := userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		utils.LogError("Failed to fetch users", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer cursor.Close(context.Background())

	var users []models.User
	if err := cursor.All(context.Background(), &users); err != nil {
		utils.LogError("Failed to decode users", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.LogError("Invalid user data", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	_, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		utils.LogError("Failed to create user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.LogError("Invalid user data", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": user}
	_, err := userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		utils.LogError("Failed to update user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	_, err := userCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		utils.LogError("Failed to delete user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
