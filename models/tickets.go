package models

import (
	"time"
)

type Ticket struct {
	Time         time.Time
	TicketString string
	Limit        int
}
