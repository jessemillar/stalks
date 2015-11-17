package helpers

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jessemillar/stalks/models"
)

func Sell(userID string, symbol string, quantity int) string {
	stock := models.CheckStock(symbol)

	// TODO: Check that users have the stock they're selling

	models.AddMoney(userID, stock.Price)

	if quantity == 0 {
		quantity = 1
	}

	models.SubtractShares(userID, symbol, quantity)
	models.AddMoney(userID, stock.Price)

	return fmt.Sprintf("%d share(s) of %s have been sold for a total of %d turnips.\n", quantity, strings.ToUpper(symbol), quantity*stock.Price) // Return information about a user's portfolio
}
