package helpers

import (
	"fmt"
	"strings"

	"github.com/jessemillar/stalks/accessors"
	"github.com/jessemillar/stalks/models"
)

// Portfolio gets the given user's portfolio and returns it in string form
func Portfolio(userID string, ag *accessors.AccessorGroup) string {
	portfolio := ag.GetPortfolio(userID)
	message := []string{}
	worth := portfolio.Turnips

	message = append(message, fmt.Sprintf("You have %s turnips in your wallet.", Comma(portfolio.Turnips)))

	for _, value := range portfolio.Investments {
		if value.Quantity > 0 {
			price := models.CheckStock(value.Ticker).Price
			worth = worth + price*value.Quantity
			message = append(message, fmt.Sprintf("You have %s share(s) of %s worth %s turnips total.", Comma(value.Quantity), value.Ticker, Comma(price*value.Quantity)))
		}
	}

	message = append(message, fmt.Sprintf("Your net worth is %s turnips.", Comma(worth)))

	response := strings.Join(message, "\n")

	return response
}
