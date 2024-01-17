package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func erro400(ctx *gin.Context, err error, message string) {
	logger.Errorf("Solicitação inválida: %s", err)
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": message,
	})
}

func erro404(ctx *gin.Context, message string) {
	logger.Infof("Entidade não encontrada")
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": message,
	})
}

func erro500(ctx *gin.Context, err error) {
	logger.Errorf("Erro no servidor: %s", err)
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "Erro no servidor",
	})
}
