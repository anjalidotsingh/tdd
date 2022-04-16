package multicurrencymoney

import (
	"fmt"
	"strconv"
	"strings"
)

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

func convertToUSD(currency string) (string, error) {
	denomination, unit, _ := strings.Cut(currency, currencyDenonimationUnitSeparator())
	denominationInFloat, _ := strconv.ParseFloat(denomination, 64)
	exchangeRate := exchangeRateForUSD(unit)
	if isExchangeRateNotAvailable(exchangeRate) {
		return "", fmt.Errorf("exchange rate for %s is not available", unit)
	}
	return fmt.Sprintf("%s USD", strconv.FormatFloat(denominationInFloat*exchangeRate, 'f', -1, 64)),nil
}

func currencyDenonimationUnitSeparator() string {
	return " "
}

func exchangeRateForUSD(currencyUnit string) float64 {
	exchangeRateForUSDMap := map[string]float64{
		"CHF": 1.5,
		"INR": 0.0125,
	}
	return exchangeRateForUSDMap[currencyUnit]
}

func isExchangeRateNotAvailable(exchangeRate float64)bool{
	return exchangeRate==0
}