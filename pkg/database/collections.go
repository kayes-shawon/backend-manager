package database

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

var DB *MongoDB

func GetServiceCollection() *mongo.Collection {
	// Assuming DB is initialized and holds the MongoDB database connection
	if DB == nil || DB.db == nil {
		panic("MongoDB connection not initialized. Call Setup() before using collections.")
	}

	collectionName := viper.GetString("SERVICE_COLLECTION_NAME")
	collection := DB.db.Collection(collectionName)
	return collection
}
