package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

type Stock struct {
	Change           float64 `json:"Change"`
	ChangePercent    float64 `json:"ChangePercent"`
	ChangePercentYTD float64 `json:"ChangePercentYTD"`
	ChangeYTD        float64 `json:"ChangeYTD"`
	High             float64 `json:"High"`
	Price            float64 `json:"LastPrice"`
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

func health(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uh, we had a slight weapons malfunction, but uh... everything's perfectly all right now. We're fine. We're all fine here now, thank you. How are you?")
}

func testPost(c web.C, w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:"+os.Getenv("STALK_DB_PASS")+"@tcp("+os.Getenv("STALK_DB")+":"+os.Getenv("STALK_DB_PORT"))

	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}

	defer db.Close()

	fmt.Fprintf(w, "%s", "Post response")
}

func price(c web.C, w http.ResponseWriter, r *http.Request) {
	client := new(http.Client)
	res, err := client.Get("http://dev.markitondemand.com/Api/v2/Quote/jsonp?symbol=" + c.URLParams["stock"])

	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	// Super ghetto method of removing unnecessary callback garbage from the API response
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}

	body = append(body[:0], body[18:]...) // Kill the leading function() bit
	body = body[:len(body)-1]             // Remove the last character so we just have naked JSON

	var stock = new(Stock)
	err = json.Unmarshal(body, stock)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}

	fmt.Fprintf(w, "%s: %f\n", stock.Name, stock.Price)
}

func main() {
	goji.Get("/health", health)
	goji.Get("/price/:stock", price)
	goji.Post("/post", testPost)
	goji.Serve()
}
