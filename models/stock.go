package models

import (
	"encoding/json"
	"log"
	"net/http"
)

type Stock struct {
	Change           float64 `json:"Change"`
	ChangePercent    float64 `json:"ChangePercent"`
	ChangePercentYTD float64 `json:"ChangePercentYTD"`
	ChangeYTD        float64 `json:"ChangeYTD"`
	High             float64 `json:"High"`
	LastPrice        float64 `json:"LastPrice"`
	Price            int
	Low              float64 `json:"Low"`
	MSDate           float64 `json:"MSDate"`
	MarketCap        int     `json:"MarketCap"`
	Name             string  `json:"Name"`
	Open             float64 `json:"Open"`
	Status           string  `json:"Status"`
	Symbol           string  `json:"Symbol"`
	Timestamp        string  `json:"Timestamp"`
	Volume           int     `json:"Volume"`
}

func CheckStock(symbol string) *Stock {
	client := new(http.Client)
	res, err := client.Get("http://dev.markitondemand.com/Api/v2/Quote/json?symbol=" + symbol)
	if err != nil { // Die if there was an error
		log.Panicf("Error: %s", err)
	}

	var stock = new(Stock)                        // Make a new instance of the Stock struct
	err = json.NewDecoder(res.Body).Decode(stock) // Populate it with our JSON data
	if err != nil {                               // Die if there was an error
		log.Panic(err)
	}

	if len(stock.Name) == 0 {
		log.Panicf("%s does not appear to be a valid stock...\n", symbol)
	}

	stock.Price = int(stock.LastPrice * 100) // Convert to turnips

	return stock
}
