package handlers

import (
	"context"
	"playlits-music/api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func getVideosPlaylistIdHandler(ctx *gin.Context, videos []models.Video) []models.VideoFull {
	var idsVideos []string
	for _, s := range videos {
		idsVideos = append(idsVideos, s.IdVideo)
	}

	collection := getDatabase().Collection("videos")

	filter := bson.M{"youtube.id": bson.M{"$in": idsVideos}}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		erro500(ctx, err)
	}

	var videosFull []models.VideoFull
	if err = cur.All(context.TODO(), &videosFull); err != nil {
		erro500(ctx, err)
	}
	defer cur.Close(context.Background())

	return videosFull
}
