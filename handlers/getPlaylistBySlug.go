package handlers

import (
	"context"
	"net/http"
	"playlits-music/api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPlaylistBySlugHandler(ctx *gin.Context) {
	slug := ctx.Param("slug")

	collection := getDatabase().Collection("playlists")

	filter := bson.M{"slug_url": slug}
	var playlist models.Playlist
	err := collection.FindOne(context.Background(), filter).Decode(&playlist)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			erro404(ctx, "Playlist não encontrada.")
		} else {
			erro500(ctx, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"item": playlist,
	})
}
