package handler

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(context *gin.Context, code int, message string) {
	context.Header("Content-Type", "application/json")
	context.JSON(code, gin.H{
		"message":   message,
		"errorcode": code,
	})
}

func sendSucess(context *gin.Context, op string, data interface{}) {
	context.Header("Content-Type", "application/json")
	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s realizado com sucesso", op),
		"data":    data,
	})
}
