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

func GetTurnips(userID string) int {
	var username string
	var turnips int

	db := accessors.Connect()

	err := db.QueryRow("SELECT * FROM users WHERE userID=?", userID).Scan(&userID, &username, &turnips)
	// if err == sql.ErrNoRows { // If the user doesn't exist yet
	// 	_, err := db.Query("INSERT INTO users (userID, username, turnips) VALUES (?,?,?)", r.PostFormValue("user_id"), r.PostFormValue("user_name"), 1000000)
	// 	if err != nil {
	// 		log.Panic("Error: %s\n", err)
	// 	}

	// 	turnips = 1000000
	// } else
	if err != nil { // Generic error
		log.Panic("Error: %s\n", err)
	}

	return turnips
}

func AddMoney(userID string, increase int) {
	db := accessors.Connect()

	turnips := GetTurnips(userID)

	_, err := db.Query("UPDATE users SET turnips=? WHERE userID=?", turnips+increase, userID)
	if err != nil { // Generic error
		log.Panic("Error: %s\n", err)
	}
}

func SubtractMoney(userID string, decrease int) {
	db := accessors.Connect()

	turnips := GetTurnips(userID)

	_, err := db.Query("UPDATE users SET turnips=? WHERE userID=?", turnips-decrease, userID)
	if err != nil { // Generic error
		log.Panic("Error: %s\n", err)
	}
}

func AddShares(userID string, stock string, quantity int) {
	db := accessors.Connect()

	_, err := db.Query("INSERT INTO portfolios (userID, ticker, quantity) VALUES (?,?,?)", userID, stock, quantity)
	if err != nil { // Generic error
		log.Panic("Error: %s\n", err)
	}
}

func SubtractShares(userID string, stock string, quantity int) {
	// TODO: Make this work
}
