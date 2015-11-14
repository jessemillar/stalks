package models

type Message struct {
	Text   string
	UserID string `json:"user_id"`
}
