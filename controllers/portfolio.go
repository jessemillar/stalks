package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/jessemillar/stalks/models"
	"github.com/zenazn/goji/web"
)

func Portfolio(c web.C, w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", os.Getenv("STALKS_DB_USER")+":"+os.Getenv("STALKS_DB_PASS")+"@tcp("+os.Getenv("STALKS_DB_HOST")+":"+os.Getenv("STALKS_DB_PORT")+")/"+os.Getenv("STALKS_DB_NAME"))
	defer db.Close()
	if err != nil { // Die if there was an error
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}

	row := db.QueryRow("SELECT * FROM portfolios WHERE userID=?", c.URLParams["userID"])
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
	}

	p := new(models.Portfolio)
	err = row.Scan(p)

	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
	}

	fmt.Fprintf(w, "%s", row)
}
