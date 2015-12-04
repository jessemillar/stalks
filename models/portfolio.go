package models

// Portfolio represents a user's portfolio complete with investments and turnips
type Portfolio struct {
	UserID      string
	Investments []Investment
	Turnips     int
}

// PortfolioValue holds user information for a whole portfolio's value
type PortfolioValue struct {
	UserID   string
	Username string
	Value    int
}

// SortedPortfolioValue is the interface used to sort portfolio values
type SortedPortfolioValue []PortfolioValue

// Len returns the lenght of the SortedPortfolioValue
func (s SortedPortfolioValue) Len() int {
	return len(s)
}

// Swap swaps positions of the SortedPortfolioValues
func (s SortedPortfolioValue) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less returns true if SortedPortfolioValues at index i is less than index j
func (s SortedPortfolioValue) Less(i, j int) bool {
	return s[i].Value < s[j].Value
}
