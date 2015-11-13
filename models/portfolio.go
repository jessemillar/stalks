package models

type Portfolio struct {
	UserID      string
	Investments []Investment
	Turnips     int
}

func (p *Portfolio) addInvestments() {
	//TODO: implement this to load database rows into the Portfolio
}
