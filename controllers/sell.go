package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zenazn/goji/web"
)

func Sell(c web.C, w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", os.Getenv("STALKS_DB_USER")+":"+os.Getenv("STALKS_DB_PASS")+"@tcp("+os.Getenv("STALKS_DB_HOST")+":"+os.Getenv("STALKS_DB_PORT")+")/"+os.Getenv("STALKS_DB_NAME"))
	defer db.Close()
	if err != nil { // Die if there was an error
		log.Printf("Error: %s\n", err)
		return
	}

	rows, err := db.Query("INSERT INTO users (firstName, lastName, userID, username, turnips) VALUES (?,?,?,?,?)", c.URLParams["firstName"], c.URLParams["lastName"], c.URLParams["username"], c.URLParams["username"], 1000000)
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
	defer rows.Close()

	log.Printf("%s", "Success")
}
