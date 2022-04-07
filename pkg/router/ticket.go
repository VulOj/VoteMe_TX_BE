package router

import (
	"VoteMe_BE_TX/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetTicket(c *gin.Context) {
	ticket := services.GetCurrentTicketDB()
	if time.Since(ticket.Time) > 2000000000 {
		services.CreateNewTicket()
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"name": ticket,
	})
	return
}
