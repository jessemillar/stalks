package helpers

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jessemillar/stalks/models"
)

func Sell(userID string, quantity int, symbol string) string {
	stock := models.CheckStock(symbol)

	holdings := models.GetShare(userID, symbol)
	if holdings >= quantity { // If we successfully sell
		models.SubtractShares(userID, symbol, quantity)
		models.AddTurnips(userID, stock.Price) // Add turnips to our wallet
	} else { // Else return a human-readable error
		return fmt.Sprintf("You do not have enough shares of %s to sell %d. You have %d shares.\n", strings.ToUpper(symbol), quantity, holdings) // Return information about a user's portfolio
	}

	return fmt.Sprintf("%d share(s) of %s have been sold for a total of %d turnips.\n", quantity, strings.ToUpper(symbol), quantity*stock.Price) // Return information about a user's portfolio
}
