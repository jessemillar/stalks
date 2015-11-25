package helpers

import (
	"fmt"
	"strings"

	"github.com/jessemillar/stalks/accessors"
	"github.com/jessemillar/stalks/models"
)

// Sell sells a given number of the user's holdings in the symbol
func Sell(userID string, quantity int, symbol string, ag *accessors.AccessorGroup) string {
	if !MarketOpen() {
		return fmt.Sprintf("The Stalk Market is currently closed.")
	}

	stock := models.CheckStock(symbol)
	user := ag.GetUser(userID)

	if stock.Name == "N/A" || stock.Price == 0 {
		return fmt.Sprintf("%s does not appear to be a valid stock...\n", strings.ToUpper(symbol)) // Return the price through the API endpoint
	}

	holdings := ag.GetShare(userID, symbol)
	if holdings >= quantity { // If we successfully sell
		ag.SubtractShares(userID, symbol, quantity)
		ag.AddTurnips(userID, stock.Price*quantity) // Add turnips to our wallet
	} else { // Else return a human-readable error
		return fmt.Sprintf("You do not have enough shares of %s to sell %s. You have %s shares.\n", strings.ToUpper(symbol), Comma(quantity), Comma(holdings)) // Return information about a user's portfolio
	}

	Webhook(fmt.Sprintf("<@%s|%s> sold %s share(s) of %s for %s turnips.", user.UserID, user.Username, Comma(quantity), strings.ToUpper(symbol), Comma(stock.Price*quantity)))

	return fmt.Sprintf("%s share(s) of %s have been sold for a total of %s turnips.\n", Comma(quantity), strings.ToUpper(symbol), Comma(quantity*stock.Price)) // Return information about a user's portfolio
}
