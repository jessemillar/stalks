package accessors

import (
	_ "github.com/go-sql-driver/mysql" // Blank import due to its use as a driver

	"database/sql"
	"log"
)

// AccessorGroup holds all configuration for the accessors.
type AccessorGroup struct {
	DB *sql.DB
}

// ConnectToDB creates a database connection and sets it in the struct
func (c *AccessorGroup) ConnectToDB(dbType, dsn string) {
	db, err := sql.Open("mysql", dsn)
	if err != nil { // Die if there was an error
		log.Panicf("Could not connect to the database: %s\n", err)
	}

	c.DB = db
}
