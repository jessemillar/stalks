package helpers

import (
	"fmt"

	"github.com/jessemillar/stalks/models"
)

func Portfolio(userID string) string {
	turnips := models.GetUser(userID).Turnips

	return fmt.Sprintf("You have %d turnips.\n", turnips)
}
