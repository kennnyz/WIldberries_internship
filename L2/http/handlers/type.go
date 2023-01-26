package handlers

type Event struct {
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
	Info   string `json:"info"`
}

var (
	Events = []Event{}
)
