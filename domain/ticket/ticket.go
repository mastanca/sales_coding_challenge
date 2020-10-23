package ticket

// Ticket - Holds ticket data such as the country and event for which it's valid
// could hold any other relevant information like creationDate, id, etc
type Ticket struct {
	Id      string `json:",omitempty" db:"id"`
	Country string `json:"country" binding:"required" db:"country"`
	Event   string `json:"event" db:"event"`
}

type Tickets []Ticket
