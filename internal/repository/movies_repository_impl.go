package repository

import (
	"context"
	"time"

	"github.com/diogocarasco/ginapi/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	client   *mongo.Client
	database string
}

func NewMongoRepository(uri, dbName string) (*MongoRepository, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		logger.Fatal("Couldn't connect to MongoDB instance", err)
		return nil, err
	}

	return &MongoRepository{client: client, database: dbName}, nil
}

func (r *MongoRepository) Close(ctx context.Context) error {
	if err := r.client.Disconnect(ctx); err != nil {
		logger.Error("Couldn't disconnect from MongoDB: ", err)
	}
	return r.client.Disconnect(ctx)
}

func (r *MongoRepository) Create(collectionName string, model interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := r.client.Database(r.database).Collection(collectionName)
	result, err := collection.InsertOne(ctx, model)

	return result, err

}

func (r *MongoRepository) Get(collectionName string, filter interface{}) (*mongo.SingleResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := r.client.Database(r.database).Collection(collectionName)
	result := collection.FindOne(ctx, filter)
	return result, result.Err()
}

func (r *Mong)
