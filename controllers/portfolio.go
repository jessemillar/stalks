package controllers

import (
	"fmt"

	"github.com/jessemillar/stalks/models"
)

func Portfolio(userID string) string {
	turnips := models.GetTurnips(userID)

	return fmt.Sprintf("You have %d turnips.\n", turnips)
}
