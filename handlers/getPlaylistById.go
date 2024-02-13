package handlers

import (
	"context"
	"net/http"
	"playlits-music/api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPlaylistByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	collection := getDatabase().Collection("playlists")

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		erro400(ctx, err, "Playlist ID inválido.")
		return
	}

	filter := bson.M{"_id": _id, "sincronizado": SINCRONIZADO}
	var playlist models.Playlist
	err = collection.FindOne(context.Background(), filter).Decode(&playlist)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			erro404(ctx, "Playlist não encontrada.")
		} else {
			erro500(ctx, err)
		}
		return
	}

	playlist.VideosCompleto = getVideosPlaylistIdHandler(ctx, playlist.Videos)
	playlist.Videos = nil

	ctx.JSON(http.StatusOK, gin.H{
		"item": playlist,
	})
}
