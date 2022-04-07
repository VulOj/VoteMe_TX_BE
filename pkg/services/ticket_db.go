package services

import (
	"VoteMe_BE_TX/models"
)

func GetCurrentTicketDB() (ticket models.Ticket) {
	db.Find(&ticket)
	return
}

func CreateTicketDB(ticket models.Ticket) {
	db.Create(&ticket)
}
func UpdateTicket(ticketstring string, limit int) {
	var ticket models.Ticket
	db.Find(&ticket)
	var temp models.Ticket
	temp.Time = ticket.Time
	temp.Limit = limit
	temp.TicketString = ticket.TicketString
	db.Delete(&ticket)
	db.Save(&temp)
}
