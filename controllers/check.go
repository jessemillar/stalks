package controllers

import (
	"fmt"
	"net/http"

	"github.com/jessemillar/stalks/helpers"
	"github.com/zenazn/goji/web"
)

func (cg *ControllerGroup) Check(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", helpers.Check(c.URLParams["symbol"]))
}
