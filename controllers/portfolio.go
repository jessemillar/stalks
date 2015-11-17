package controllers

import (
	"fmt"
	"net/http"

	"github.com/jessemillar/stalks/models"
	"github.com/zenazn/goji/web"
)

func Portfolio(c web.C, w http.ResponseWriter, r *http.Request) {
	turnips := models.GetTurnips(r.PostFormValue("user_id"))

	fmt.Fprintf(w, "You have %d turnips.\n", turnips) // Return information about a user's portfolio
}
