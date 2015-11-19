package controllers

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

// User creates a given user account
func (cg *ControllerGroup) User(c web.C, w http.ResponseWriter, r *http.Request) {
	user := cg.Accessors.GetUser(r.PostFormValue("user_id"))
	if len(user.Username) > 0 {
		fmt.Fprintf(w, "Your account already exists. You have %d turnips.\n", user.Turnips)
	} else {
		cg.Accessors.MakeUser(r.PostFormValue("user_id"), r.PostFormValue("user_name"))
		user = cg.Accessors.GetUser(r.PostFormValue("user_id"))

		fmt.Fprintf(w, "Your account has been created and supplied with %d turnips. Welcome to the Stalk Exchange!\n", user.Turnips)
	}
}
