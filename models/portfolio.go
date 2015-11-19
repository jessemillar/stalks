package models

// Portfolio represents a user's portfolio complete with investments and turnips
type Portfolio struct {
	UserID      string
	Investments []Investment
	Turnips     int
}
