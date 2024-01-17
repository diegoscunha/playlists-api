package handlers

import (
	"context"
	"math"
	"net/http"
	"playlits-music/api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPlaylistsHandler(ctx *gin.Context) {
	// tratamento query string
	limit, err := strconv.ParseInt(ctx.DefaultQuery("limit", strconv.Itoa(int(RESULTS_PER_PAGE))), 10, 64)
	if err != nil {
		limit = RESULTS_PER_PAGE
	}
	page, err := strconv.ParseInt(ctx.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		page = 1
	}

	skip := limit * (page - 1)

	//collection := db.Database("playlists_db").Collection("playlists")
	collection := getDatabase().Collection("playlists")

	filter := bson.M{"sincronizado": SINCRONIZADO}

	totalResults, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		erro500(ctx, err)
	}

	ops := options.Find().
		SetLimit(limit).
		SetSkip(skip).
		SetProjection(bson.D{
			{Key: "_id", Value: 1},
			{Key: "data_criacao", Value: 1},
			{Key: "youtube.id", Value: 1},
			{Key: "youtube.data_criacao", Value: 1},
			{Key: "youtube.id_canal", Value: 1},
			{Key: "youtube.titulo_canal", Value: 1},
			{Key: "youtube.titulo", Value: 1},
			{Key: "youtube.descricao", Value: 1},
			{Key: "youtube.numero_videos", Value: 1},
			{Key: "youtube.imagens", Value: 1},
		})
	cur, err := collection.Find(context.Background(), filter, ops)
	if err != nil {
		erro500(ctx, err)
	}

	var playlists []models.Playlist
	if err = cur.All(context.TODO(), &playlists); err != nil {
		erro500(ctx, err)
	}
	defer cur.Close(context.Background())

	// tratamento paginas
	var numPages = math.Round(float64(totalResults) / float64(limit))
	var nextPage int64 = 0
	if int64(numPages) > page {
		nextPage = page + 1
	}
	var prevPage int64 = 0
	if page > 1 {
		prevPage = page - 1
	}

	if playlists == nil {
		playlists = []models.Playlist{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"nextPage": nextPage,
		"prevPage": prevPage,
		"pageInfo": gin.H{
			"totalResults":   totalResults,
			"resultsPerPage": limit,
		},
		"items": playlists,
	})
}
