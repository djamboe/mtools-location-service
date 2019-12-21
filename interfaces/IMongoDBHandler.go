package interfaces

import (
	"github.com/djamboe/mtools-location-service/models"
	"go.mongodb.org/mongo-driver/bson"
)

type IMongoDBHandler interface {
	FindOne(filter bson.M, collectionName string, databaseName string) (IRowMongoDB, error)
	InsertOne(param models.Point, collectionName string, databaseName string) error
}
type IRowMongoDB interface {
	Next() bool
	DecodeResults(v interface{}) error
}
