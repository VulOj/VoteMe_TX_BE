package services

import (
	"VoteMe_BE_TX/models"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func CreateNewTicket() {
	rand.Seed(time.Now().UnixNano())
	oneTicket := models.Ticket{
		TicketString: CreatRandomTicketString(10),
		Time:         time.Now(),
		Limit:        10,
	}
	CreateTicketDB(oneTicket)
}

func CreatRandomTicketString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
