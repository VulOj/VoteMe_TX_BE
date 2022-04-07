package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "access-control-allow-origin", "token"},
		AllowCredentials: false,
	}
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	api := r.Group("/api")
	{
		user := api.Group("/auth")
		user.Use(cors.Default())
		{
			user.GET("/hello", HelloWorld)
			user.GET("/ticket", GetTicket)
			user.GET("/candidate-list", GetCandidates)
			user.GET("/candidate", GetOneCandidate)
			user.POST("/vote", Vote)
		}
	}
}
