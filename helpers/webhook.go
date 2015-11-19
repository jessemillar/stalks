package helpers

import (
	"log"
	"os"

	"github.com/parnurzeal/gorequest"
)

func Webhook(message string) {
	request := gorequest.New()
	resp, body, err := request.Post(os.Getenv("STALKS_SLACK_WEBHOOK")).
		Send(`{"text":"` + message + `"}`).
		End()
	if err != nil { // Die if there was an error
		log.Panic(err)
	}

	log.Println(resp)
	log.Println(body)
}
