package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloWorld(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Hello World!",
	})
	return
}
