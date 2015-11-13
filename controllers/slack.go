package controllers

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

func Slack(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Println(c) // "Log" what we get from Slack
	fmt.Println(r)
}
