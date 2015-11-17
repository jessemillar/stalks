package accessors

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("STALKS_DB_USER")+":"+os.Getenv("STALKS_DB_PASS")+"@tcp("+os.Getenv("STALKS_DB_HOST")+":"+os.Getenv("STALKS_DB_PORT")+")/"+os.Getenv("STALKS_DB_NAME"))
	if err != nil { // Die if there was an error
		log.Printf("Error: %s\n", err)
		return nil
	}

	return db
}
