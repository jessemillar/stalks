package controllers

import (
	"net/http"
	"strings"

	"github.com/zenazn/goji/web"
)

func Slack(c web.C, w http.ResponseWriter, r *http.Request) {
	params := strings.Fields(r.PostFormValue("text"))

	if params[0] == "check" {
		Check(c, w, r)
	} else if params[0] == "portfolio" {
		Portfolio(c, w, r)
	} else if params[0] == "buy" {
		Buy(c, w, r)
	}
}
