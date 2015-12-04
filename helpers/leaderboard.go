package helpers

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jessemillar/stalks/accessors"
	"github.com/jessemillar/stalks/models"
)

// ReportLeaders returns a string of the leaderboard
func ReportLeaders(ag *accessors.AccessorGroup) string {
	users := ag.GetAllUsers()
	pValues := []models.PortfolioValue{}

	// Compile portfolio data
	for _, user := range users {
		portfolio := ag.GetPortfolio(user.UserID)
		worth := portfolio.Turnips

		for _, value := range portfolio.Investments {
			if value.Quantity > 0 {
				price := models.CheckStock(value.Ticker).Price
				worth = worth + price*value.Quantity
			}
		}

		pValues = append(pValues, models.PortfolioValue{UserID: user.UserID, Username: user.Username, Value: worth})

	}

	// Sort the portfolios by value
	sort.Sort(models.SortedPortfolioValue(pValues))

	message := []string{}
	message = append(message, fmt.Sprintf("*End of the Day Leaderboard*"))
	// Run through the sorted values and compile the message
	for _, pValue := range pValues {
		message = append(message, fmt.Sprintf("<@%s|%s> has a net worth of %s turnips.", pValue.UserID, pValue.Username, Comma(pValue.Value)))
	}

	response := strings.Join(message, "\\n") // Double escape the newline because Slack incoming webhooks are obsessive with JSON formatting while the /slash-command "endpoints" are now

	return response
}
