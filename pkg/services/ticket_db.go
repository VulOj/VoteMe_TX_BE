package services

import "VoteMe_BE_TX/models"

func GetCurrentTicketDB() (ticket models.Ticket) {
	db.Find(&ticket)
	return
}

func CreateTicketDB(ticket models.Ticket) {
	db.Create(&ticket)
}
