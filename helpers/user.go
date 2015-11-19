package helpers

import (
	"fmt"

	"github.com/jessemillar/stalks/accessors"
)

// MakeUser creates the given user and puts them into the database
func MakeUser(userID string, username string, ag *accessors.AccessorGroup) string {
	user := ag.GetUser(userID)
	if len(user.Username) > 0 {
		return fmt.Sprintf("Your account already exists. You have %s turnips.\n", Comma(user.Turnips))
	}

	ag.MakeUser(userID, username)
	user = ag.GetUser(userID)

	return fmt.Sprintf("Your account has been created and supplied with %s turnips. Welcome to the Stalk Exchange!\n", Comma(user.Turnips))
}
