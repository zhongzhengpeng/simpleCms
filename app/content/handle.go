package content

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GET(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK, gin.H{
			"message": "get",
		})
}

func POST(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK, gin.H{
			"message": "post",
		})
}

func PUT(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK, gin.H{
			"message": "put",
		})
}

func DELETE(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "delete",
	})
}
