package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zenazn/goji/web"
)

func Portfolio(c web.C, w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", os.Getenv("STALKS_DB_USER")+":"+os.Getenv("STALKS_DB_PASS")+"@tcp("+os.Getenv("STALKS_DB_HOST")+":"+os.Getenv("STALKS_DB_PORT")+")/"+os.Getenv("STALKS_DB_NAME"))
	defer db.Close()
	if err != nil { // Die if there was an error
		log.Printf("Error: %s\n", err)
		return
	}

	var userID, username string
	var turnips int
	err = db.QueryRow("SELECT * FROM users WHERE userID=?", r.PostFormValue("user_id")).Scan(&userID, &username, &turnips)
	if err == sql.ErrNoRows { // If the user doesn't exist yet
		_, err := db.Query("INSERT INTO users (userID, username, turnips) VALUES (?,?,?)", r.PostFormValue("user_id"), r.PostFormValue("user_name"), 1000000)
		if err != nil {
			log.Printf("Error: %s\n", err)
		}

		fmt.Fprintf(w, "You have %d turnips.\n", 1000000) // Return information about a user's portfolio
		return
	} else if err != nil { // Generic error
		log.Printf("Error: %s\n", err)
		return
	}

	fmt.Fprintf(w, "You have %d turnips.\n", turnips) // Return information about a user's portfolio
}
