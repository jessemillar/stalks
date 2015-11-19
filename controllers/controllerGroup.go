package controllers

import (
	"database/sql"
	"log"
)

// ControllerGroup holds all config information for the controllers
type ControllerGroup struct {
	DB *sql.DB
}

// ConnectToDB creates a database connection and sets it in the struct
func (c *ControllerGroup) ConnectToDB(dbType, dsn string) {
	db, err := sql.Open("mysql", dsn)
	if err != nil { // Die if there was an error
		log.Panicf("Could not connect to the database: %s\n", err)
	}

	c.DB = db
}
