package main

import (
	"os"

	"github.com/jessemillar/stalks/controllers"
	"github.com/zenazn/goji"
)

func main() {

	// Construct the dsn used for the database
	dsn := os.Getenv("STALKS_DB_USER") + ":" + os.Getenv("STALKS_DB_PASS") + "@tcp(" + os.Getenv("STALKS_DB_HOST") + ":" + os.Getenv("STALKS_DB_PORT") + ")/" + os.Getenv("STALKS_DB_NAME")

	// Construct a new controllerGroup and connect to the database
	cg := new(controllers.ControllerGroup)
	cg.ConnectToDB("mysql", dsn)

	goji.Get("/health", cg.Health)
	goji.Post("/slack", cg.Slack) // The main endpoint that Slack hits
	goji.Post("/play", cg.User)
	goji.Post("/portfolio", cg.Portfolio)
	goji.Get("/check/:symbol", cg.Check)
	goji.Post("/buy/:quantity/:symbol", cg.Buy)
	goji.Post("/sell/:quantity/:symbol", cg.Sell)
	goji.Serve()
}
