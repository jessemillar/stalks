package models

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Stock struct {
	Name  string
	Ask   float64
	Bid   float64
	Price int
}

func CheckStock(symbol string) *Stock {
	client := new(http.Client)
	res, err := client.Get("http://finance.yahoo.com/d/quotes.csv?s=" + symbol + "&f=nab")
	if err != nil {
		log.Panic(err)
	}

	var stock = new(Stock) // Make a new instance of the Stock struct

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}

	response := strings.Split(string(body), ",")

	ask, _ := strconv.ParseFloat(response[1], 64)
	bid, _ := strconv.ParseFloat(response[2], 64)

	stock.Name = response[0]
	stock.Ask = ask * 100
	stock.Bid = bid * 100

	if stock.Ask > stock.Bid { // Ghetto and temporary
		stock.Price = int(stock.Ask)
	} else {
		stock.Price = int(stock.Bid)
	}

	return stock
}
