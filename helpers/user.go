package helpers

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jessemillar/stalks/models"
)

func MakeUser(userID string, username string) string {
	user := models.GetUser(userID)
	if len(user.Username) > 0 {
		return fmt.Sprintf("Your account already exists. You have %d turnips.\n", user.Turnips)
	} else {
		models.MakeUser(userID, username)
		user = models.GetUser(userID)

		return fmt.Sprintf("Your account has been created and supplied with %d turnips. Welcome to the Stalk Exchange!\n", user.Turnips)
	}
}
