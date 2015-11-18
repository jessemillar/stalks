package main

import (
	"github.com/jessemillar/stalks/controllers"
	"github.com/zenazn/goji"
)

func main() {
	goji.Get("/health", controllers.Health)
	goji.Post("/slack", controllers.Slack) // The main endpoint that Slack hits
	goji.Post("/play", controllers.User)
	goji.Post("/portfolio", controllers.Portfolio)
	goji.Get("/check/:symbol", controllers.Check)
	goji.Post("/buy/:symbol/:quantity", controllers.Buy)
	goji.Post("/sell/:symbol/:quantity", controllers.Sell)
	goji.Serve()
}
