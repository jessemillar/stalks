package helpers

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jessemillar/stalks/models"
)

func Buy(userID string, quantity int, symbol string) string {
	stock := models.CheckStock(symbol)
	turnips := models.GetUser(userID).Turnips

	// Make sure they have enough turnips to buy
	if turnips < stock.Price*quantity {
		return fmt.Sprintf("%d shares of %s costs %s turnips and you have %s turnips.\n", quantity, strings.ToUpper(symbol), Comma(stock.Price*quantity), Comma(turnips)) // Return information about a user's portfolio
	}

	models.SubtractTurnips(userID, stock.Price*quantity)

	if quantity == 0 {
		quantity = 1
	}

	models.AddShares(userID, symbol, quantity)

	return fmt.Sprintf("%s turnips were spent to add %d share(s) of %s to your portfolio.\n", Comma(stock.Price*quantity), quantity, strings.ToUpper(symbol)) // Return information about a user's portfolio
}
