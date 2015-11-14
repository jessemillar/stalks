package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jessemillar/stalks/models"
	"github.com/zenazn/goji/web"
)

func Check(c web.C, w http.ResponseWriter, r *http.Request) {
	client := new(http.Client)
	res, err := client.Get("http://dev.markitondemand.com/Api/v2/Quote/json?symbol=" + c.URLParams["stock"])
	if err != nil { // Die if there was an error
		log.Printf("Error: %s", err)
		return
	}

	var stock = new(models.Stock)                 // Make a new instance of the Stock struct
	err = json.NewDecoder(res.Body).Decode(stock) // Populate it with our JSON data
	if err != nil {                               // Die if there was an error
		log.Printf("Error: %s\n", err)
		return
	}

	if len(stock.Name) > 0 {
		fmt.Fprintf(w, "{%s is currently worth %d turnips.\n", stock.Name, int(stock.Price*100)) // Return the price through the API endpoint
	} else {
		fmt.Fprintf(w, "%s does not appear to be a valid stock...\n", c.URLParams["stock"]) // Return the price through the API endpoint
	}
}
