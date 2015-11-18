package models

import (
	"database/sql"
	"log"

	"github.com/jessemillar/stalks/accessors"
)

type Holding struct {
	ID       int
	UserID   string
	Ticker   string
	Quantity int
}

type Portfolio struct {
	UserID      string
	Investments []Investment
	Turnips     int
}

func GetShares(userID string, symbol string) int {
	db := accessors.Connect()

	var holding = new(Holding)
	err := db.QueryRow("SELECT * FROM portfolios WHERE userID=? AND ticker=?", userID, symbol).Scan(&holding.ID, &holding.UserID, &holding.Ticker, &holding.Quantity)
	if err == sql.ErrNoRows { // If the user doesn't have any holdings of the given stock
		return -1
	} else if err != nil {
		log.Panic(err)
	}

	return holding.Quantity
}

func AddShares(userID string, symbol string, increase int) {
	db := accessors.Connect()

	quantity := GetShares(userID, symbol)
	if quantity >= 0 {
		_, err := db.Query("UPDATE portfolios SET quantity=? WHERE userID=? AND ticker=?", quantity+increase, userID, symbol)
		if err != nil {
			log.Panic(err)
		}
	} else {
		_, err := db.Query("INSERT INTO portfolios (userID, ticker, quantity) VALUES (?,?,?)", userID, symbol, increase)
		if err != nil {
			log.Panic(err)
		}
	}
}

func SubtractShares(userID string, symbol string, decrease int) int {
	db := accessors.Connect()

	quantity := GetShares(userID, symbol)
	if quantity >= 0 && quantity >= decrease {
		_, err := db.Query("UPDATE portfolios SET quantity=? WHERE userID=? AND ticker=?", quantity-decrease, userID, symbol)
		if err != nil {
			log.Panic(err)
		}

		return quantity - decrease
	} else { // You didn't have enough holdings
		return quantity // Return the current number of holdings
	}
}
