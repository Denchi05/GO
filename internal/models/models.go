package models

type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	ImageURL string `json:"imageUrl"`
}

type BookingRequest struct {
	SessionID string `json:"sessionId"`
	Row       int    `json:"row"`
	Seat      int    `json:"seat"`
}

type BookingResponse struct {
	Success  bool   `json:"success"`
	TicketID string `json:"ticketId"`
	Message  string `json:"message"`
}
