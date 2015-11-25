package models

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Stock struct {
	Name  string
	Price int
}

func CheckStock(symbol string) *Stock {
	client := new(http.Client)
	res, err := client.Get("http://finance.yahoo.com/d/quotes.csv?s=" + symbol + "&f=l1on")
	if err != nil {
		log.Panic(err)
	}

	var stock = new(Stock) // Make a new instance of the Stock struct

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}

	response := strings.Split(string(body), ",")

	price, err := strconv.ParseFloat(response[0], 64)
	if err != nil {
		log.Panic(err)
	}

	for i := 2; i < len(response); i++ { // Takes care of stupid stocks with names like "Groupon, Inc." that throw off the strings.Split() method above
		stock.Name = fmt.Sprint(stock.Name, response[i])
	}

	stock.Name = strings.TrimSuffix(stock.Name, "\n") // Kill the silly newline character

	stock.Price = int(price * 100) // Convert to turnips

	return stock
}
