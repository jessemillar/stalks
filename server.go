package main

import (
	"github.com/jessemillar/stalks/controllers"
	"github.com/zenazn/goji"
)

func main() {
	// goji.Get("/health", controllers.Health)
	goji.Post("/slack", controllers.Slack) // The main endpoint that Slack hits
	// goji.Post("/portfolio", controllers.Portfolio)
	// goji.Get("/check/:stock", controllers.Check)
	// goji.Post("/buy/:stock/:quantity", controllers.Buy)
	// goji.Post("/sell/:stock/:quantity", controllers.Sell)
	goji.Serve()
}
