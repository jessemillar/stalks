package helpers

import (
	"fmt"
	"strings"

	"github.com/jessemillar/stalks/models"
)

func Buy(userID string, quantity int, symbol string) string {
	stock := models.CheckStock(symbol)
	user := models.GetUser(userID)

	// Make sure they have enough turnips to buy
	if user.Turnips < stock.Price*quantity {
		return fmt.Sprintf("%s shares of %s costs %s turnips and you have %s turnips.\n", Comma(quantity), strings.ToUpper(symbol), Comma(stock.Price*quantity), Comma(user.Turnips)) // Return information about a user's portfolio
	}

	models.SubtractTurnips(userID, stock.Price*quantity)

	if quantity == 0 {
		quantity = 1
	}

	models.AddShares(userID, symbol, quantity)

	Webhook(fmt.Sprintf("<@%s|%s> purchased %s share(s) of %s for %s turnips.", user.UserID, user.Username, Comma(quantity), strings.ToUpper(symbol), Comma(stock.Price*quantity)))

	return fmt.Sprintf("%s turnips were spent to add %s share(s) of %s to your portfolio.\n", Comma(stock.Price*quantity), Comma(quantity), strings.ToUpper(symbol)) // Return information about a user's portfolio
}
