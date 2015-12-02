package helpers

import (
	"fmt"
	"strings"

	"github.com/jessemillar/stalks/accessors"
	"github.com/jessemillar/stalks/models"
)

func ReportLeaders(ag *accessors.AccessorGroup) string {
	message := []string{}
	users := ag.GetAllUsers()

	message = append(message, fmt.Sprintf("*End of the Day Leaderboard*"))

	for _, user := range users {
		portfolio := ag.GetPortfolio(user.UserID)
		worth := portfolio.Turnips

		for _, value := range portfolio.Investments {
			if value.Quantity > 0 {
				price := models.CheckStock(value.Ticker).Price
				worth = worth + price*value.Quantity
			}
		}

		message = append(message, fmt.Sprintf("<@%s|%s>'s net worth is %s turnips.", user.UserID, user.Username, Comma(worth)))
	}

	response := strings.Join(message, "\n")

	return response
}
