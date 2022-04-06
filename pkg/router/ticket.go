package router

import (
	"VoteMe_BE_TX/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetTicket(c *gin.Context) {
	ticket := services.GetCurrentTicketDB()
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":  time.Since(ticket.Time),
		"name": ticket,
	})
	return
}
