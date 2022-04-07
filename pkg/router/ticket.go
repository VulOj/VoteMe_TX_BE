package router

import (
	"VoteMe_BE_TX/pkg/consts"
	"VoteMe_BE_TX/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetTicket(c *gin.Context) {
	ticket := services.GetCurrentTicketDB()
	if time.Since(ticket.Time) > consts.VALID_TIME {
		services.CreateNewTicket()
	}
	c.JSON(http.StatusOK, gin.H{
		"name": ticket,
	})
	return
}
