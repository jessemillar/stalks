package accessors

import (
	"database/sql"
	"log"

	"github.com/jessemillar/stalks/models"
)

// MakeUser adds a user to the database
func (ag *AccessorGroup) MakeUser(userID string, username string) {
	_, err := ag.DB.Query("INSERT INTO users (userID, username, turnips) VALUES (?,?,?)", userID, username, 1000000)
	if err != nil {
		log.Panic(err)
	}
}

// GetUser returns a user from the database by userID
func (ag *AccessorGroup) GetUser(userID string) *models.User {
	var user = new(models.User)
	err := ag.DB.QueryRow("SELECT * FROM users WHERE userID=?", userID).Scan(&user.UserID, &user.Username, &user.Turnips)

	if err == sql.ErrNoRows { // If the user doesn't exist yet
		return user // Return a blank user
	} else if err != nil {
		log.Panic(err)
	}

	return user
}

// GetAllUsers returns an array of all users from the database
func (ag *AccessorGroup) GetAllUsers() []*models.User {
	rows, err := ag.DB.Query("SELECT * FROM users")

	if err != nil {
		log.Panic(err)
	}

	var users []*models.User

	for rows.Next() {
		var newUser = new(models.User)
		err = rows.Scan(&newUser.UserID, &newUser.Username, &newUser.Turnips)
		if err != nil {
			log.Panic(err)
		}

		users = append(users, newUser)
	}

	return users
}

// AddTurnips adds a given number of turnips to the given userID
func (ag *AccessorGroup) AddTurnips(userID string, increase int) {
	turnips := ag.GetUser(userID).Turnips

	_, err := ag.DB.Query("UPDATE users SET turnips=? WHERE userID=?", turnips+increase, userID)
	if err != nil {
		log.Panic(err)
	}
}

// SubtractTurnips removes a given number of turnips from the given user
func (ag *AccessorGroup) SubtractTurnips(userID string, decrease int) {
	turnips := ag.GetUser(userID).Turnips

	_, err := ag.DB.Query("UPDATE users SET turnips=? WHERE userID=?", turnips-decrease, userID)
	if err != nil {
		log.Panic(err)
	}
}
