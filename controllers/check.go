package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jessemillar/stalks/models"
	"github.com/zenazn/goji/web"
)

func Check(c web.C, w http.ResponseWriter, r *http.Request) {
	params := strings.Fields(r.PostFormValue("text"))

	client := new(http.Client)
	res, err := client.Get("http://dev.markitondemand.com/Api/v2/Quote/json?symbol=" + params[1])
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
		fmt.Fprintf(w, "%s is currently worth %d turnips.\n", stock.Name, int(stock.Price*100)) // Return the price through the API endpoint
	} else {
		fmt.Fprintf(w, "%s does not appear to be a valid stock...\n", params[1]) // Return the price through the API endpoint
	}
}
