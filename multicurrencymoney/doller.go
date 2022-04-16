package multicurrencymoney

import (
	"fmt"
	"strconv"
	"strings"
)

type Bond struct {
	Instrument string
	Shares     int
	Price      string
}

func calculateTotalPrices(bonds []Bond) (string, error) {
	var totalBondsPrice float64
	for _, bond := range bonds {
		denomination, unit, _ := strings.Cut(bond.Price, currencyDenonimationUnitSeparator())
		if unit != "USD" {
			temp, err := convertToUSD(bond.Price)
			if err != nil {
				return "", fmt.Errorf("exchange rate for %s is not available", unit)
			}
			denomination, _, _ = strings.Cut(temp, currencyDenonimationUnitSeparator())
		}
		denominationInFloat, _ := strconv.ParseFloat(denomination, 64)
		totalBondsPrice += denominationInFloat * float64(bond.Shares)
	}
	return fmt.Sprintf("%s USD", strconv.FormatFloat(totalBondsPrice, 'f', -1, 64)), nil
}

func convertToUSD(currency string) (string, error) {
	denomination, unit, _ := strings.Cut(currency, currencyDenonimationUnitSeparator())
	denominationInFloat, _ := strconv.ParseFloat(denomination, 64)
	exchangeRate := exchangeRateForUSD(unit)
	if isExchangeRateNotAvailable(exchangeRate) {
		return "", fmt.Errorf("exchange rate for %s is not available", unit)
	}
	return fmt.Sprintf("%s USD", strconv.FormatFloat(denominationInFloat*exchangeRate, 'f', -1, 64)), nil
}

func currencyDenonimationUnitSeparator() string {
	return " "
}

func exchangeRateForUSD(currencyUnit string) float64 {
	exchangeRateForUSDMap := map[string]float64{
		"CHF": 1.5,
		"INR": 0.0125,
		"USD": 1,
	}
	return exchangeRateForUSDMap[currencyUnit]
}

func isExchangeRateNotAvailable(exchangeRate float64) bool {
	return exchangeRate == 0
}
