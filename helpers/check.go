package helpers

import (
	"fmt"
	"strings"

	"github.com/jessemillar/stalks/models"
)

// Check checks the given symbol for current price
func Check(symbol string) string {
	stock := models.CheckStock(symbol)

	if len(stock.Name) > 0 {
		return fmt.Sprintf("%s (%s) is currently worth %s turnips.\n", stock.Name, strings.ToUpper(symbol), Comma(stock.Price)) // Return the price through the API endpoint
	}

	return fmt.Sprintf("%s does not appear to be a valid stock...\n", strings.ToUpper(symbol)) // Return the price through the API endpoint
}
