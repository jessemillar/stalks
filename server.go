package main

import (
	"os"

	"github.com/jessemillar/stalks/accessors"
	"github.com/jessemillar/stalks/controllers"
	"github.com/jessemillar/stalks/helpers"
	"github.com/robfig/cron"
	"github.com/zenazn/goji"
)

func main() {
	// Construct the dsn used for the database
	dsn := os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") + ")/" + os.Getenv("DATABASE_NAME")

	// Construct a new AccessorGroup and connects it to the database
	ag := new(accessors.AccessorGroup)
	ag.ConnectToDB("mysql", dsn)

	// Constructs a new ControllerGroup and gives it the AccessorGroup
	cg := new(controllers.ControllerGroup)
	cg.Accessors = ag

	c := cron.New()
	c.AddFunc("0 0 21 * * 1-5", func() { // Run at 2:00pm MST (which is 21:00 UTC) Monday through Friday
		helpers.Webhook(helpers.ReportLeaders(ag))
	})
	c.Start()

	goji.Get("/health", cg.Health)
	goji.Get("/leaderboard", cg.ReportLeaders)
	goji.Post("/slack", cg.Slack) // The main endpoint that Slack hits
	goji.Post("/play", cg.User)
	goji.Post("/portfolio", cg.Portfolio)
	goji.Get("/check/:symbol", cg.Check)
	goji.Post("/buy/:quantity/:symbol", cg.Buy)
	goji.Post("/sell/:quantity/:symbol", cg.Sell)
	goji.Serve()
}
