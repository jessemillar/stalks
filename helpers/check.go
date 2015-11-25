package helpers

import (
	"fmt"
	"strings"

	"github.com/jessemillar/stalks/models"
)

// Check checks the given symbol for current price
func Check(symbol string) string {
	stock := models.CheckStock(symbol)

	if stock.Name != "N/A" && stock.Price > 0 { // Make sure Yahoo gave us a valid stock
		return fmt.Sprintf("%s (%s) is currently worth %s turnips (%s).\n", stock.Name, strings.ToUpper(symbol), Comma(stock.Price), stock.Change) // Return the price through the API endpoint
	} else {
		return fmt.Sprintf("%s does not appear to be a valid stock...\n", strings.ToUpper(symbol)) // Return the price through the API endpoint
	}
}
