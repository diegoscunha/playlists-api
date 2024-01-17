package handlers

import (
	"playlits-music/api/configs"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	logger *configs.Logger
	db     *mongo.Client
	config *configs.Config
)

const RESULTS_PER_PAGE int64 = 2
const SINCRONIZADO string = "S"

func InitHandler() {
	logger = configs.GetLogger("handlers")
	db = configs.GetMongodb()
	config = configs.GetConfig()
}

func getDatabase() *mongo.Database {
	return db.Database(config.MongoDbName)
}
