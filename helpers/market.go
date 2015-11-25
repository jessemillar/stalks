package helpers

func MarketOpen() bool {
	// now := time.Now()
	// utc, err := time.LoadLocation("America/New_York")
	// if err != nil {
	// 	log.Panic(err)
	// }

	// now = now.In(utc) // Convert the current, local time to the timezone returned by the API

	// timestamp, err := time.Parse("Mon Jan 2 15:04:05 UTC-05:00 2006", models.CheckStock("AAPL").Timestamp)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// if now.Minute() == timestamp.Minute() && now.Hour() == timestamp.Hour() && now.Day() == timestamp.Day() {
	// 	log.Println("true")
	// 	return true
	// } else {
	// 	log.Println("false")
	// 	return false
	// }

	return true
}
