package controllers

import (
	"context"
	"crm-system/api/models"
	"crm-system/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var authCollection *mongo.Collection

func init() {
	authCollection = utils.GetMongoClient().Database("crm").Collection("auth")
}

func Login(c *gin.Context) {
	var credentials models.Auth
	if err := c.ShouldBindJSON(&credentials); err != nil {
		utils.LogError("Invalid login data", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login data"})
		return
	}

	var user models.Auth
	err := authCollection.FindOne(context.Background(), bson.M{"username": credentials.Username}).Decode(&user)
	if err != nil || !utils.CheckPasswordHash(credentials.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func Register(c *gin.Context) {
	var newUser models.Auth
	if err := c.ShouldBindJSON(&newUser); err != nil {
		utils.LogError("Invalid registration data", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid registration data"})
		return
	}

	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		utils.LogError("Failed to hash password", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
		return
	}

	newUser.Password = hashedPassword
	_, err = authCollection.InsertOne(context.Background(), newUser)
	if err != nil {
		utils.LogError("Failed to create user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}
