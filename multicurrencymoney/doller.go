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

func CalculateTotalPrices(bonds []Bond) (string, error) {
	var totalBondsPrice float64
	for _, bond := range bonds {
		denomination, err := convertCurrencyToDefaultUnit(bond.Price)
		if err != nil {
			return "", err
		}
		totalBondsPrice += denomination * float64(bond.Shares)
	}
	return fmt.Sprintf("%s %s", strconv.FormatFloat(totalBondsPrice, 'f', -1, 64), getDefaultCurrencyUnit()), nil
}

func convertCurrencyToDefaultUnit(currency string) (float64, error) {
	denominationInString, unit, _ := strings.Cut(currency, currencyDenonimationUnitSeparator())
	denomination, err := strconv.ParseFloat(denominationInString, 64)
	if unit == getDefaultCurrencyUnit() {
		return denomination, err
	}
	denomination, err = ConvertToUSD(denomination, unit)
	if err != nil {
		return denomination, err
	}
	return denomination, err
}

func ConvertToUSD(denomination float64, unit string) (float64, error) {
	exchangeRate := exchangeRateForUSD(unit)
	if !isExchangeRateAvailable(exchangeRate) {
		return 0, fmt.Errorf("exchange rate for %s is not available", unit)
	}
	return denomination * exchangeRate, nil
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

func isExchangeRateAvailable(exchangeRate float64) bool {
	return exchangeRate != 0
}

func getDefaultCurrencyUnit() string {
	return "USD"
}
