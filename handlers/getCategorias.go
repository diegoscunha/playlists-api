package handlers

import (
	"context"
	"net/http"
	"playlits-music/api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCategoriasHandler(ctx *gin.Context) {
	collection := getDatabase().Collection("categorias")

	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		erro500(ctx, err)
	}

	var categorias []models.Categoria
	if err = cur.All(context.TODO(), &categorias); err != nil {
		erro500(ctx, err)
	}
	defer cur.Close(context.Background())

	ctx.JSON(http.StatusOK, gin.H{
		"items": categorias,
	})
}
