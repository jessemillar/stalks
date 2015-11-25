package models

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Stock struct {
	Name   string
	Price  int
	Change string
}

func CheckStock(symbol string) *Stock {
	client := new(http.Client)
	res, err := client.Get("http://finance.yahoo.com/d/quotes.csv?s=" + symbol + "&f=l1p2n")
	if err != nil {
		log.Panic(err)
	}

	var stock = new(Stock) // Make a new instance of the Stock struct

	reader := csv.NewReader(res.Body)
	response, err := reader.Read()
	if err != nil {
		log.Panic(err)
	}

	// Ghetto "error" out
	if response[0] == "N/A" { // This is Yahoo's way of telling us that the stock we're looking for is invalid
		stock.Name = "N/A"
		stock.Price = 0
		stock.Change = "0%"

		return stock
	}

	price, err := strconv.ParseFloat(response[0], 64)
	if err != nil {
		log.Panic(err)
	}

	for i := 2; i < len(response); i++ { // Takes care of stupid stocks with names like "Groupon, Inc." that throw off the strings.Split() method above
		stock.Name = fmt.Sprint(stock.Name, response[i])
	}

	stock.Name = strings.TrimSuffix(stock.Name, "\n") // Kill the silly newline character that Yahoo returns

	stock.Price = int(price * 100) // Convert to turnips
	stock.Change = response[1]

	return stock
}
