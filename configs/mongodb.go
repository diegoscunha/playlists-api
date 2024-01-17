package configs

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongodb() (*mongo.Client, error) {
	logger := GetLogger("mongodb")
	config := GetConfig()

	connect := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", config.MongoUser, config.MongoPassword, config.MongoHost, config.MongoPort, config.MongoDbName)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connect))
	if err != nil {
		logger.Errorf("falha conexão mongodb: %v", err)
		return nil, err
	}

	return client, nil
}
