package database

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseManager struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewDatabaseManager(uri, databaseName, collectionName string) *DatabaseManager {
	client := ConnectDB(uri)
	collection := newCollection(client, databaseName, collectionName)
	return &DatabaseManager{client, collection}
}

func (dbm *DatabaseManager) Close() {
	dbm.client.Disconnect(context.Background())
}

func (dbm *DatabaseManager) GetClient() *mongo.Client {
	return dbm.client
}

func (dbm *DatabaseManager) GetCollection() *mongo.Collection {
	return dbm.collection
}

func (dbm *DatabaseManager) SetCollection(databaseName, collectionName string) {
	dbm.collection = newCollection(dbm.client, databaseName, collectionName)
}

// ConnectDB connects to the database
func ConnectDB(uri string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// GetCollection returns a collection from the database
func newCollection(client *mongo.Client, databaseName, collectionName string) *mongo.Collection {
	collection := client.
		Database(databaseName).
		Collection(collectionName)
	return collection
}

// InsertOne inserts a document into the database
func InsertOne(collection *mongo.Collection, document interface{}) (string, error) {
	found := collection.FindOne(context.Background(), document)
	if found.Err() == nil {
		return "", errors.New("document already exists")
	}
	if found.Err() != mongo.ErrNoDocuments {
		return "", found.Err()
	}
	res, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func DeleteById(collection *mongo.Collection, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(context.Background(), primitive.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}
