package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jessemillar/stalks/helpers"
	"github.com/zenazn/goji/web"
)

func Slack(c web.C, w http.ResponseWriter, r *http.Request) {
	params := strings.Fields(r.PostFormValue("text"))

	if params[0] == "play" {
		fmt.Fprintf(w, "%s\n", helpers.MakeUser(r.PostFormValue("user_id"), r.PostFormValue("user_name")))
	} else if params[0] == "check" || params[0] == "c" {
		fmt.Fprintf(w, "%s\n", helpers.Check(params[1]))
	} else if params[0] == "portfolio" || params[0] == "p" {
		fmt.Fprintf(w, "%s\n", helpers.Portfolio(r.PostFormValue("user_id")))
	} else if params[0] == "buy" || params[0] == "b" {
		quantity, _ := strconv.Atoi(params[1])
		fmt.Fprintf(w, "%s\n", helpers.Buy(r.PostFormValue("user_id"), quantity, params[2]))
	} else if params[0] == "sell" || params[0] == "s" {
		quantity, _ := strconv.Atoi(params[1])
		fmt.Fprintf(w, "%s\n", helpers.Sell(r.PostFormValue("user_id"), quantity, params[2]))
	}
}
