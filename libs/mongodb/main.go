package mongodb

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type CollectionName string

var (
	TaskCollection CollectionName = "tasks"
	UserCollection CollectionName = "users"
)

type IMongoDb interface {
	InsertOne(ctx context.Context, collectionName CollectionName, T interface{}) (*mongo.InsertOneResult, error)
	GetOne(ctx context.Context, collectionName CollectionName, filter interface{}) (*mongo.SingleResult, error)
	UpdateOne(ctx context.Context, collectionName CollectionName, filter, update interface{}) (*mongo.UpdateResult, error)
}

type mongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

func Initialize(uri, databaseName string) IMongoDb {
	log.Info().Msgf("Database %v %v", uri, databaseName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().
		ApplyURI(uri).SetConnectTimeout(10 * time.Second).
		SetServerSelectionTimeout(10 * time.Second)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Panic().Err(err).Send()
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Panic().Err(err).Send()
	}
	db := client.Database(databaseName)

	// Tạo unique index
	uuidIndexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "uuid", Value: 1}}, // Tạo index trên trường "uuid"
		Options: options.Index().SetUnique(true), // Đặt unique = true
	}
	emailIndexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}}, // Tạo index trên trường "email"
		Options: options.Index().SetUnique(true),  // Đặt unique = true
	}

	db.Collection(string(TaskCollection)).Indexes().CreateOne(ctx, uuidIndexModel)

	db.Collection(string(UserCollection)).Indexes().CreateMany(ctx, []mongo.IndexModel{
		uuidIndexModel,
		emailIndexModel,
	})

	return &mongoDB{
		client: client,
		db:     db,
	}
}

func (db *mongoDB) InsertOne(ctx context.Context, collectionName CollectionName, T interface{}) (*mongo.InsertOneResult, error) {
	collection := db.db.Collection(string(collectionName))
	return collection.InsertOne(ctx, T)
}

func (db *mongoDB) GetOne(ctx context.Context, collectionName CollectionName, filter interface{}) (*mongo.SingleResult, error) {
	collection := db.db.Collection(string(collectionName))
	rs := collection.FindOne(ctx, filter)
	if rs.Err() != nil {
		return nil, rs.Err()
	}
	return rs, nil
}

func (db *mongoDB) UpdateOne(ctx context.Context, collectionName CollectionName, filter, update interface{}) (*mongo.UpdateResult, error) {
	collection := db.db.Collection(string(collectionName))
	return collection.UpdateOne(ctx, filter, update)
}
