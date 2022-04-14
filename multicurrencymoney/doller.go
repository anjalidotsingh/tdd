package multicurrencymoney

type Bond struct {
	Instrument string
	Shares     int
	Price      float64
}

func calculateTotalPrices(bonds []Bond) float64 {
	var totalBondsPrice float64
	for _, bond := range bonds {
		totalBondsPrice += bond.Price * float64(bond.Shares)
	}
	return totalBondsPrice
}
