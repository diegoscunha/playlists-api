package router

import (
	"playlits-music/api/configs"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() {
	config := configs.GetConfig()
	// Inicialização do router
	router := gin.Default()
	router.Use(initCors())
	router.NoRoute(notFound)
	// Inicialização das rotas
	initRoutes(router)
	// Rodando api
	router.Run(":" + config.ServerPort)
}

func initCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:9000"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	})
}

func notFound(ctx *gin.Context) {
	ctx.JSON(404, gin.H{"message": "Page not found"})
}
