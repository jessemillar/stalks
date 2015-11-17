package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/jessemillar/stalks/models"
	"github.com/zenazn/goji/web"
)

func Check(c web.C, w http.ResponseWriter, r *http.Request) {
	params := strings.Fields(r.PostFormValue("text"))

	stock := models.CheckStock(params[1])

	if len(stock.Name) > 0 {
		fmt.Fprintf(w, "%s is currently worth %d turnips.\n", stock.Name, stock.Price) // Return the price through the API endpoint
	} else {
		fmt.Fprintf(w, "%s does not appear to be a valid stock...\n", params[1]) // Return the price through the API endpoint
	}
}
