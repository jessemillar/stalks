package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jessemillar/stalks/models"
	"github.com/zenazn/goji/web"
)

func Buy(c web.C, w http.ResponseWriter, r *http.Request) {
	params := strings.Fields(r.PostFormValue("text"))

	stock := models.CheckStock(params[1])
	turnips := models.GetTurnips(r.PostFormValue("user_id"))

	// Make sure they have enough turnips to buy
	if turnips < int(stock.Price) {
		fmt.Fprintf(w, "You do not have enough turnips.\n") // Return information about a user's portfolio
		return
	}

	models.SubtractMoney(r.PostFormValue("user_id"), stock.Price)

	quantity, _ := strconv.Atoi(params[2])

	if quantity == 0 {
		quantity = 1
	}

	models.AddShares(r.PostFormValue("user_id"), params[1], quantity)

	fmt.Fprintf(w, "%d share(s) of %s have been added to your portfolio.\n", quantity, params[1]) // Return information about a user's portfolio
}
