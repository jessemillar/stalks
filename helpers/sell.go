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
		models.AddTurnips(userID, stock.Price*quantity) // Add turnips to our wallet
	} else { // Else return a human-readable error
		return fmt.Sprintf("You do not have enough shares of %s to sell %s. You have %s shares.\n", strings.ToUpper(symbol), Comma(quantity), Comma(holdings)) // Return information about a user's portfolio
	}

	Webhook(fmt.Sprintf("<@%s|%s> sold %s share(s) of %s for %s turnips.", user.UserID, user.Username, Comma(quantity), strings.ToUpper(symbol), Comma(stock.Price*quantity)))

	return fmt.Sprintf("%s share(s) of %s have been sold for a total of %s turnips.\n", Comma(quantity), strings.ToUpper(symbol), Comma(quantity*stock.Price)) // Return information about a user's portfolio
}
