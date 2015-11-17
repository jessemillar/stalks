package models

import (
	"log"

	"github.com/jessemillar/stalks/accessors"
)

type Portfolio struct {
	UserID      string
	Investments []Investment
	Turnips     int
}

func AddShares(userID string, stock string, quantity int) {
	db := accessors.Connect()

	_, err := db.Query("INSERT INTO portfolios (userID, ticker, quantity) VALUES (?,?,?)", userID, stock, quantity)
	if err != nil { // Generic error
		log.Panic(err)
	}
}

func SubtractShares(userID string, stock string, quantity int) {
	// TODO: Make this work
}
