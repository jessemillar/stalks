package helpers

import (
	"log"
	"os"

	"github.com/parnurzeal/gorequest"
)

// Webhook calls out to the configured webhook and sends a message
func Webhook(message string) {
	endpoint := os.Getenv("SLACK_WEBHOOK")

	if len(endpoint) == 0 { // If the webhook endpoint is not set, ignore
		return
	}

	request := gorequest.New()
	resp, body, err := request.Post(os.Getenv("SLACK_WEBHOOK")).
		Send(`{"text":"` + message + `"}`).
		End()
	if err != nil { // Die if there was an error
		log.Panic(err)
	}

	// Leave these log statements here in case Slack starts being dumb
	log.Println(resp)
	log.Println(body)
}
