package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName         = "world"
	humansCollectionName = "humans"
)

type MongoDB interface {
	GetMongoDBClient() *mongo.Client
	GetHumansCollection() *mongo.Collection
	CloseConnection(ctx context.Context)
}

type mongoDB struct {
	Client           *mongo.Client
	HumansCollection *mongo.Collection
}

func NewMongoDB() (MongoDB, error) {
	var err error
	clientOptions := options.Client().ApplyURI(os.Getenv("CONNECT_URI"))
	mongoDBClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return &mongoDB{}, err
	}
	humansCollection := mongoDBClient.Database(databaseName).Collection(humansCollectionName)
	return &mongoDB{Client: mongoDBClient, HumansCollection: humansCollection}, err
}

func (m *mongoDB) GetMongoDBClient() *mongo.Client {
	return m.Client
}

func (m *mongoDB) CloseConnection(ctx context.Context) {
	log.Panicln(m.Client.Disconnect(ctx))
}

func (m *mongoDB) GetHumansCollection() *mongo.Collection {
	return m.HumansCollection
}
