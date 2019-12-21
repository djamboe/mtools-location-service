package infrastructures

import (
	"context"
	"fmt"
	"github.com/djamboe/mtools-location-service/interfaces"
	"github.com/djamboe/mtools-location-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBHandler struct {
	Conn *mongo.Client
}

func (handler *MongoDBHandler) FindOne(filter bson.M, collectionName string, dbName string) (interfaces.IRowMongoDB, error) {
	collection := handler.Conn.Database(dbName).Collection(collectionName)
	rows := collection.FindOne(context.TODO(), filter)
	row := new(MongoRow)
	row.Rows = rows
	return row, nil
}

func (handler *MongoDBHandler) InsertOne(data models.Point, collectionName string, dbName string) error {
	collection := handler.Conn.Database(dbName).Collection(collectionName)
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

type MongoRow struct {
	Rows *mongo.SingleResult
}

func (r *MongoRow) DecodeResults(v interface{}) error {
	return r.Rows.Decode(v)
}

func (r MongoRow) Next() bool {
	return r.Next()
}
