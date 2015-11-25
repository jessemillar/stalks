package helpers

import (
	"fmt"
	"strings"

	"github.com/jessemillar/stalks/accessors"
	"github.com/jessemillar/stalks/models"
)

// Buy buys a given number of stocks for the user
func Buy(userID string, quantity int, symbol string, ag *accessors.AccessorGroup) string {
	if !MarketOpen() {
		return fmt.Sprintf("The Stalk Market is currently closed.")
	}

	stock := models.CheckStock(symbol)
	user := ag.GetUser(userID)

	if stock.Name == "N/A" || stock.Price == 0 {
		return fmt.Sprintf("%s does not appear to be a valid stock... Please try again.\n", strings.ToUpper(symbol)) // Return the price through the API endpoint
	}

	// Make sure they have enough turnips to buy
	if user.Turnips < stock.Price*quantity {
		return fmt.Sprintf("%s shares of %s costs %s turnips and you have %s turnips.\n", Comma(quantity), strings.ToUpper(symbol), Comma(stock.Price*quantity), Comma(user.Turnips)) // Return information about a user's portfolio
	}

	ag.SubtractTurnips(userID, stock.Price*quantity)

	if quantity == 0 {
		quantity = 1
	}

	ag.AddShares(userID, symbol, quantity)

	Webhook(fmt.Sprintf("<@%s|%s> purchased %s share(s) of %s for %s turnips.", user.UserID, user.Username, Comma(quantity), strings.ToUpper(symbol), Comma(stock.Price*quantity)))

	return fmt.Sprintf("%s turnips were spent to add %s share(s) of %s to your portfolio.\n", Comma(stock.Price*quantity), Comma(quantity), strings.ToUpper(symbol)) // Return information about a user's portfolio
}
