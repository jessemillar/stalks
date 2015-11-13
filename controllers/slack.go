package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jessemillar/stalks/models"
	"github.com/zenazn/goji/web"
)

func Slack(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Println(c) // "Log" what we get from Slack
	fmt.Println(r)

	var message = new(models.Message) // Make a new instance of the Stock struct

	err := json.NewDecoder(r.Body).Decode(message) // Populate it with our JSON data
	if err != nil {                                // Die if there was an error
		log.Printf("Error: %s\n", err)
		return
	}
}
