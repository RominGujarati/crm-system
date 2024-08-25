package utils

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func GetMongoClient() *mongo.Client {
	return client
}
