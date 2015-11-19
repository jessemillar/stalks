package helpers

import (
	"log"
	"os"

	"github.com/parnurzeal/gorequest"
)

// Webhook calls out to the configured webhook and sends a message
func Webhook(message string) {
	endpoint := os.Getenv("STALKS_SLACK_WEBHOOK")

	if len(endpoint) == 0 { // if the webhook endpoint is not set ignore
		return
	}

	request := gorequest.New()
	resp, body, err := request.Post(os.Getenv("STALKS_SLACK_WEBHOOK")).
		Send(`{"text":"` + message + `"}`).
		End()
	if err != nil { // Die if there was an error
		log.Panicf("Error: %s", err)
	}

	log.Println(resp)
	log.Println(body)
}
