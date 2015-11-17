package helpers

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jessemillar/stalks/models"
)

func Buy(userID string, symbol string, quantity int) string {
	stock := models.CheckStock(symbol)
	turnips := models.GetTurnips(userID)

	// Make sure they have enough turnips to buy
	if turnips < int(stock.Price) {
		return fmt.Sprintf("You do not have enough turnips.\n") // Return information about a user's portfolio
	}

	models.SubtractMoney(userID, stock.Price)

	if quantity == 0 {
		quantity = 1
	}

	models.AddShares(userID, symbol, quantity)

	return fmt.Sprintf("%d share(s) of %s have been added to your portfolio.\n", quantity, strings.ToUpper(symbol)) // Return information about a user's portfolio
}
