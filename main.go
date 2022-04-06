package main

import (
	"VoteMe_BE_TX/pkg/router"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8080")
}
