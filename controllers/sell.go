package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jessemillar/stalks/helpers"
	"github.com/zenazn/goji/web"
)

func Sell(c web.C, w http.ResponseWriter, r *http.Request) {
	quantity, _ := strconv.Atoi(c.URLParams["quantity"])
	fmt.Fprintf(w, "%s\n", helpers.Sell(r.PostFormValue("userID"), quantity, c.URLParams["symbol"]))
}
