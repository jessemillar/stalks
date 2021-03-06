package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jessemillar/stalks/helpers"
	"github.com/zenazn/goji/web"
)

// Buy handles http requests to purchace a given stock
func (cg *ControllerGroup) Buy(c web.C, w http.ResponseWriter, r *http.Request) {
	quantity, _ := strconv.Atoi(c.URLParams["quantity"])
	fmt.Fprintf(w, "%s\n", helpers.Buy(r.PostFormValue("userID"), quantity, c.URLParams["symbol"], cg.Accessors))
}
