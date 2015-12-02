package main

import (
	"log"
	"os"

	"github.com/jessemillar/stalks/accessors"
	"github.com/jessemillar/stalks/controllers"
	"github.com/jessemillar/stalks/helpers"
	"github.com/robfig/cron"
	"github.com/zenazn/goji"
)

func main() {
	// Construct the dsn used for the database
	dsn := os.Getenv("STALKS_DB_USER") + ":" + os.Getenv("STALKS_DB_PASS") + "@tcp(" + os.Getenv("STALKS_DB_HOST") + ":" + os.Getenv("STALKS_DB_PORT") + ")/" + os.Getenv("STALKS_DB_NAME")

	// Construct a new AccessorGroup and connects it to the database
	ag := new(accessors.AccessorGroup)
	ag.ConnectToDB("mysql", dsn)

	// Constructs a new ControllerGroup and gives it the AccessorGroup
	cg := new(controllers.ControllerGroup)
	cg.Accessors = ag

	c := cron.New()
	// c.AddFunc("0 0 18 * * 1-5", func() { helpers.Webhook(helpers.ReportLeaders(ag)) }) // Run at 4pm MST Monday through Friday
	c.AddFunc("0 8 15 * * 1-5", func() {
		log.Println("Running")
		log.Println(helpers.ReportLeaders(ag))
	})
	c.Start()

	goji.Get("/health", cg.Health)
	goji.Get("/leaders", cg.ReportLeaders)
	goji.Post("/slack", cg.Slack) // The main endpoint that Slack hits
	goji.Post("/play", cg.User)
	goji.Post("/portfolio", cg.Portfolio)
	goji.Get("/check/:symbol", cg.Check)
	goji.Post("/buy/:quantity/:symbol", cg.Buy)
	goji.Post("/sell/:quantity/:symbol", cg.Sell)
	goji.Serve()
}
