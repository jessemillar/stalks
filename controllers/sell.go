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

func Sell(c web.C, w http.ResponseWriter, r *http.Request) {
	params := strings.Fields(r.PostFormValue("text"))

	stock := models.CheckStock(params[1])

	// TODO: Check that users have the stock they're selling

	models.AddMoney(r.PostFormValue("user_id"), stock.Price)

	quantity, _ := strconv.Atoi(params[2])

	if quantity == 0 {
		quantity = 1
	}

	models.SubtractShares(r.PostFormValue("user_id"), params[1], quantity)
	models.AddMoney(r.PostFormValue("user_id"), stock.Price)

	fmt.Fprintf(w, "%d share(s) of %s have been sold for a total of %d turnips.\n", quantity, params[1], quantity*stock.Price) // Return information about a user's portfolio
}
