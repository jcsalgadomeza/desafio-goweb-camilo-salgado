package domains

type Ticket struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Destination string `json:"destination"`
	Date        string `json:"date"`
}
