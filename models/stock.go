package models

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
