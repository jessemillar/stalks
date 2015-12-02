package accessors

import (
	"database/sql"
	"log"
	"strings"

	"github.com/jessemillar/stalks/models"
)

// GetShare gets the given user's investment in a given symbol
func (ag *AccessorGroup) GetShare(userID string, symbol string) int {
	var investment = new(models.Investment)
	err := ag.DB.QueryRow("SELECT * FROM portfolios WHERE userID=? AND ticker=?", userID, strings.ToUpper(symbol)).Scan(&investment.ID, &investment.UserID, &investment.Ticker, &investment.Quantity)
	if err == sql.ErrNoRows { // If the user doesn't have any shares of the given stock
		return -1
	} else if err != nil {
		log.Panic(err)
	}

	return investment.Quantity
}

// GetAllShares returns all of the given user's investments
func (ag *AccessorGroup) GetAllShares(userID string) []models.Investment {
	var investments []models.Investment
	rows, err := ag.DB.Query("SELECT * FROM portfolios WHERE userID=?", userID)
	if err != nil {
		log.Panic(err)
	}

	for rows.Next() {
		var share = new(models.Investment)
		rows.Scan(&share.ID, &share.UserID, &share.Ticker, &share.Quantity)

		investments = append(investments, *share)
	}

	return investments
}

// GetPortfolio gets the given user's portfolio
func (ag *AccessorGroup) GetPortfolio(userID string) models.Portfolio {
	var portfolio = new(models.Portfolio)

	portfolio.UserID = userID
	portfolio.Turnips = ag.GetUser(userID).Turnips
	portfolio.Investments = ag.GetAllShares(userID)

	return *portfolio
}

// AddShares adds a number of shares of a symbol to the given user
func (ag *AccessorGroup) AddShares(userID string, symbol string, increase int) {
	quantity := ag.GetShare(userID, symbol)
	if quantity >= 0 {
		_, err := ag.DB.Query("UPDATE portfolios SET quantity=? WHERE userID=? AND ticker=?", quantity+increase, userID, strings.ToUpper(symbol))
		if err != nil {
			log.Panic(err)
		}
	} else {
		_, err := ag.DB.Query("INSERT INTO portfolios (userID, ticker, quantity) VALUES (?,?,?)", userID, strings.ToUpper(symbol), increase)
		if err != nil {
			log.Panic(err)
		}
	}
}

// SubtractShares removes a certain number of given shares from the user and returns the new quantity
func (ag *AccessorGroup) SubtractShares(userID string, symbol string, decrease int) int {
	quantity := ag.GetShare(userID, symbol)
	if quantity >= 0 && quantity >= decrease {
		_, err := ag.DB.Query("UPDATE portfolios SET quantity=? WHERE userID=? AND ticker=?", quantity-decrease, userID, strings.ToUpper(symbol))
		if err != nil {
			log.Panic(err)
		}

		return quantity - decrease
	}

	return quantity // Return the current number of holdings

}
