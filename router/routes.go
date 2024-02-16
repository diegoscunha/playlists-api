package router

import (
	"net/http"
	"playlits-music/api/handlers"

	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {
	handlers.InitHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/info", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Api rodando",
			})
		})

		v1.GET("/playlists", handlers.GetPlaylistsHandler)
		v1.GET("/playlists/search", handlers.GetPlaylistsSearchHandler)
		v1.GET("/playlists/:id", handlers.GetPlaylistByIdHandler)
		v1.GET("/playlists/slug/:slug", handlers.GetPlaylistBySlugHandler)
		v1.GET("/playlists/:id/videos", handlers.GetVideosByPlaylistIdHandler)

		v1.GET("/categorias", handlers.GetCategoriasHandler)
	}
}
