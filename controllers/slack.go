package controllers

import (
	"net/http"
	"strings"

	"github.com/zenazn/goji/web"
)

func Slack(c web.C, w http.ResponseWriter, r *http.Request) {
	params := strings.Fields(r.PostFormValue("text"))

	if params[0] == "check" || params[0] == "c" {
		Check(c, w, r)
	} else if params[0] == "portfolio" || params[0] == "p" {
		Portfolio(c, w, r)
	} else if params[0] == "buy" || params[0] == "b" {
		Buy(c, w, r)
	} else if params[0] == "sell" || params[0] == "s" {
		Sell(c, w, r)
	}
}
