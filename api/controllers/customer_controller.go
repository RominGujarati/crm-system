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

var customerCollection *mongo.Collection

func init() {
	customerCollection = utils.GetMongoClient().Database("crm").Collection("customers")
}

func GetCustomers(c *gin.Context) {
	cursor, err := customerCollection.Find(context.Background(), bson.M{})
	if err != nil {
		utils.LogError("Failed to fetch customers", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch customers"})
		return
	}
	defer cursor.Close(context.Background())

	var customers []models.Customer
	if err := cursor.All(context.Background(), &customers); err != nil {
		utils.LogError("Failed to decode customers", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode customers"})
		return
	}

	c.JSON(http.StatusOK, customers)
}

func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		utils.LogError("Invalid customer data", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer data"})
		return
	}

	_, err := customerCollection.InsertOne(context.Background(), customer)
	if err != nil {
		utils.LogError("Failed to create customer", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		utils.LogError("Invalid customer data", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer data"})
		return
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": customer}
	_, err := customerCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		utils.LogError("Failed to update customer", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	_, err := customerCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		utils.LogError("Failed to delete customer", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted"})
}
