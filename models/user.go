package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jessemillar/stalks/accessors"
)

type User struct {
	UserID   string
	Username string
	Turnips  int
}

func MakeUser(userID string, username string) {
	db := accessors.Connect()

	_, err := db.Query("INSERT INTO users (userID, username, turnips) VALUES (?,?,?)", userID, username, 1000000)
	if err != nil {
		log.Panic(err)
	}
}

func GetUser(userID string) *User {
	db := accessors.Connect()

	var user = new(User)
	err := db.QueryRow("SELECT * FROM users WHERE userID=?", userID).Scan(&user.UserID, &user.Username, &user.Turnips)
	if err == sql.ErrNoRows { // If the user doesn't exist yet
		return user // Return a blank user
	} else if err != nil {
		log.Panic(err)
	}

	return user
}

func AddTurnips(userID string, increase int) {
	db := accessors.Connect()

	turnips := GetUser(userID).Turnips

	_, err := db.Query("UPDATE users SET turnips=? WHERE userID=?", turnips+increase, userID)
	if err != nil {
		log.Panic(err)
	}
}

func SubtractTurnips(userID string, decrease int) {
	db := accessors.Connect()

	turnips := GetUser(userID).Turnips

	_, err := db.Query("UPDATE users SET turnips=? WHERE userID=?", turnips-decrease, userID)
	if err != nil {
		log.Panic(err)
	}
}
