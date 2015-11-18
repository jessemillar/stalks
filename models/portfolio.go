package models

import (
	"database/sql"
	"log"
	"strings"

	"github.com/jessemillar/stalks/accessors"
)

type Portfolio struct {
	UserID      string
	Investments []Investment
	Turnips     int
}

func GetShare(userID string, symbol string) int {
	db := accessors.Connect()

	var investment = new(Investment)
	err := db.QueryRow("SELECT * FROM portfolios WHERE userID=? AND ticker=?", userID, strings.ToUpper(symbol)).Scan(&investment.ID, &investment.UserID, &investment.Ticker, &investment.Quantity)
	if err == sql.ErrNoRows { // If the user doesn't have any shares of the given stock
		return -1
	} else if err != nil {
		log.Panic(err)
	}

	return investment.Quantity
}

func GetAllShares(userID string) []Investment {
	db := accessors.Connect()

	var investments []Investment
	rows, err := db.Query("SELECT * FROM portfolios WHERE userID=?", userID)
	if err != nil {
		log.Panic(err)
	}

	for rows.Next() {
		var share = new(Investment)
		rows.Scan(&share.ID, &share.UserID, &share.Ticker, &share.Quantity)

		investments = append(investments, *share)
	}

	return investments
}

func GetPortfolio(userID string) Portfolio {
	var portfolio = new(Portfolio)

	portfolio.UserID = userID
	portfolio.Turnips = GetUser(userID).Turnips
	portfolio.Investments = GetAllShares(userID)

	return *portfolio
}

func AddShares(userID string, symbol string, increase int) {
	db := accessors.Connect()

	quantity := GetShare(userID, symbol)
	if quantity >= 0 {
		_, err := db.Query("UPDATE portfolios SET quantity=? WHERE userID=? AND ticker=?", quantity+increase, userID, strings.ToUpper(symbol))
		if err != nil {
			log.Panic(err)
		}
	} else {
		_, err := db.Query("INSERT INTO portfolios (userID, ticker, quantity) VALUES (?,?,?)", userID, strings.ToUpper(symbol), increase)
		if err != nil {
			log.Panic(err)
		}
	}
}

func SubtractShares(userID string, symbol string, decrease int) int {
	db := accessors.Connect()

	quantity := GetShare(userID, symbol)
	if quantity >= 0 && quantity >= decrease {
		_, err := db.Query("UPDATE portfolios SET quantity=? WHERE userID=? AND ticker=?", quantity-decrease, userID, strings.ToUpper(symbol))
		if err != nil {
			log.Panic(err)
		}

		return quantity - decrease
	} else { // You didn't have enough holdings
		return quantity // Return the current number of holdings
	}
}
