package helpers

import (
	"fmt"
	"strings"

	"github.com/jessemillar/stalks/accessors"
	"github.com/jessemillar/stalks/models"
)

// Buy buys a given number of stocks for the user
func Buy(userID string, quantity int, symbol string, ag *accessors.AccessorGroup) string {
	stock := models.CheckStock(symbol)
	turnips := ag.GetUser(userID).Turnips

	// Make sure they have enough turnips to buy
	if turnips < stock.Price*quantity {
		return fmt.Sprintf("%d shares of %s costs %s turnips and you have %s turnips.\n", quantity, strings.ToUpper(symbol), Comma(stock.Price*quantity), Comma(turnips)) // Return information about a user's portfolio
	}

	ag.SubtractTurnips(userID, stock.Price*quantity)

	if quantity == 0 {
		quantity = 1
	}

	ag.AddShares(userID, symbol, quantity)

	return fmt.Sprintf("%s turnips were spent to add %d share(s) of %s to your portfolio.\n", Comma(stock.Price*quantity), quantity, strings.ToUpper(symbol)) // Return information about a user's portfolio
}
