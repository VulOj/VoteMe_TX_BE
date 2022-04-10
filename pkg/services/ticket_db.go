package services

import (
	"VoteMe_BE_TX/models"
)

func GetCurrentTicketDB() (ticket models.Ticket) {
	db.Last(&ticket)
	return
}

func CreateTicketDB(ticket models.Ticket) {
	db.Create(&ticket)
}
func DeleteOldTicket() {
	var ticket models.Ticket
	db.Delete(&ticket)
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
