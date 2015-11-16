package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jessemillar/stalks/models"
	"github.com/zenazn/goji/web"
)

func Buy(c web.C, w http.ResponseWriter, r *http.Request) {
	params := strings.Fields(r.PostFormValue("text"))

	db, err := sql.Open("mysql", os.Getenv("STALKS_DB_USER")+":"+os.Getenv("STALKS_DB_PASS")+"@tcp("+os.Getenv("STALKS_DB_HOST")+":"+os.Getenv("STALKS_DB_PORT")+")/"+os.Getenv("STALKS_DB_NAME"))
	defer db.Close()
	if err != nil { // Die if there was an error
		log.Printf("Error: %s\n", err)
		return
	}

	client := new(http.Client)
	res, err := client.Get("http://dev.markitondemand.com/Api/v2/Quote/json?symbol=" + params[1])
	if err != nil { // Die if there was an error
		log.Printf("Error: %s", err)
		return
	}

	var stock = new(models.Stock)                 // Make a new instance of the Stock struct
	err = json.NewDecoder(res.Body).Decode(stock) // Populate it with our JSON data
	if err != nil {                               // Die if there was an error
		log.Printf("Error: %s\n", err)
		return
	}

	if len(stock.Name) == 0 {
		fmt.Fprintf(w, "%s does not appear to be a valid stock...\n", params[1]) // Return an error
		return
	}

	// Get total turnips for this user
	var userID, username string
	var turnips int
	err = db.QueryRow("SELECT * FROM users WHERE userID=?", r.PostFormValue("user_id")).Scan(&userID, &username, &turnips)
	if err == sql.ErrNoRows { // If the user doesn't exist yet
		_, err := db.Query("INSERT INTO users (userID, username, turnips) VALUES (?,?,?)", r.PostFormValue("user_id"), r.PostFormValue("user_name"), 1000000)
		if err != nil {
			log.Printf("Error: %s\n", err)
		}

		turnips = 1000000
		fmt.Fprintf(w, "You have %d turnips.\n", 1000000) // Return information about a user's portfolio
		return
	} else if err != nil { // Generic error
		log.Printf("Error: %s\n", err)
		return
	}

	// Make sure they have enough turnips to buy
	if turnips < int(stock.Price) {
		fmt.Fprintf(w, "You do not have enough turnips.\n") // Return information about a user's portfolio
		return
	}

	// Subtract money
	_, err = db.Query("UPDATE users SET turnips=? WHERE userID=?", float64(turnips)-stock.Price, r.PostFormValue("user_id"))
	if err != nil { // Generic error
		log.Printf("Error: %s\n", err)
		return
	}

	// Convert the string params[2] to an int
	var quantity int
	_, err = fmt.Sscan(params[2], &quantity)

	if quantity == 0 {
		quantity = 1
	}

	_, err = db.Query("INSERT INTO portfolios (userID, ticker, quantity) VALUES (?,?,?)", r.PostFormValue("user_id"), params[1], quantity)
	if err != nil { // Generic error
		log.Printf("Error: %s\n", err)
		return
	}

	fmt.Fprintf(w, "%d share(s) of %s have been added to your portfolio.\n", quantity, params[1]) // Return information about a user's portfolio
}
