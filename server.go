package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jessemillar/stalks/controllers"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func makeUser(c web.C, w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", os.Getenv("STALKS_DB_USER")+":"+os.Getenv("STALKS_DB_PASS")+"@tcp("+os.Getenv("STALKS_DB_HOST")+":"+os.Getenv("STALKS_DB_PORT")+")/"+os.Getenv("STALKS_DB_NAME"))
	defer db.Close()
	if err != nil { // Die if there was an error
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}

	rows, err := db.Query("INSERT INTO users (firstName, lastName, userID, username, turnips) VALUES (?,?,?,?,?)", c.URLParams["firstName"], c.URLParams["lastName"], c.URLParams["username"], c.URLParams["username"], 1000000)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
	}
	defer rows.Close()

	fmt.Fprintf(w, "%s", "Success")
}

func main() {
	goji.Get("/health", controllers.Health)
	goji.Post("/slack", controllers.Slack) // The main endpoint that Slack hits
	goji.Get("/portfolio/:userID", controllers.Portfolio)
	goji.Get("/check/:stock", controllers.Check)
	goji.Post("/makeUser/:username/:firstName/:lastName", makeUser) // Mostly for development purposes
	// goji.Post("/buy/:stock/:quantity", buy)
	// goji.Post("/sell/:stock/:quantity", sell)
	goji.Serve()
}
