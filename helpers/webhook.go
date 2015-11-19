package helpers

import (
	"bytes"
	"net/http"
	"os"
)

func Webhook(message string) {
	url := os.Getenv("STALKS_SLACK_WEBHOOK")

	var jsonStr = []byte(`{"payload":"Buy cheese and bread for breakfast."}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
}
