package ticket

// Ticket - Holds ticket data such as the country and event for which it's valid
// could hold any other relevant information like creationDate, id, etc
type Ticket struct {
	Country string `json:"country" binding:"required"`
	Event   string `json:"event"`
}

type Tickets []Ticket
