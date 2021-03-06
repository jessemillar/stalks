package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jessemillar/stalks/helpers"
	"github.com/zenazn/goji/web"
)

// Slack is a special kind of handler, it handles all requests from slack and routes them accordingly
func (cg *ControllerGroup) Slack(c web.C, w http.ResponseWriter, r *http.Request) {
	params := strings.Fields(r.PostFormValue("text"))

	if len(params) == 0 {
		fmt.Fprintf(w, "%s\n", "Please enter a command. Usage: `/stalk [play, portfolio, check, buy, sell]`")
		return
	}

	if params[0] == "play" {
		fmt.Fprintf(w, "%s\n", helpers.MakeUser(r.PostFormValue("user_id"), r.PostFormValue("user_name"), cg.Accessors))
	}

	user := cg.Accessors.GetUser(r.PostFormValue("user_id"))
	if len(user.Username) == 0 { // If we get a blank user returned
		fmt.Fprintf(w, "Your account does not exist yet. Please run `/stalk play` to start.")
		return
	}

	if params[0] == "check" || params[0] == "c" {
		if len(params) < 2 {
			fmt.Fprintf(w, "Not enough parameters. Usage: `/stalk check [symbol]`")
			return
		}

		fmt.Fprintf(w, "%s\n", helpers.Check(params[1]))
	} else if params[0] == "portfolio" || params[0] == "p" {
		fmt.Fprintf(w, "%s\n", helpers.Portfolio(r.PostFormValue("user_id"), cg.Accessors))
	} else if params[0] == "buy" || params[0] == "b" {
		if len(params) < 3 {
			fmt.Fprintf(w, "Not enough parameters. Usage: `/stalk buy [quantity] [symbol]`")
			return
		}

		quantity, _ := strconv.Atoi(params[1])
		fmt.Fprintf(w, "%s\n", helpers.Buy(r.PostFormValue("user_id"), quantity, params[2], cg.Accessors))
	} else if params[0] == "sell" || params[0] == "s" {
		if len(params) < 2 {
			fmt.Fprintf(w, "Not enough parameters. Usage: `/stalk sell [quantity] [symbol]`")
			return
		}

		quantity, _ := strconv.Atoi(params[1])
		fmt.Fprintf(w, "%s\n", helpers.Sell(r.PostFormValue("user_id"), quantity, params[2], cg.Accessors))
	}
}
