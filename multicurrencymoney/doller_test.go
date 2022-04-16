package multicurrencymoney

import (
	"testing"
)

/*
to-do list
1-add calculate price method for USD
2-add exchange rate to USD
-for CHF
-for other currency
-for cuurency that doesnt exist in exchange table
3-add calculate price support to USD for multiple currency using exchange rate
*/

func TestCalculateTotalPrice(t *testing.T) {
	type args struct {
		bonds []Bond
	}
	tests := []struct {
		name           string
		args           args
		wantTotalPrice float64
	}{
		{
			name:           "Instrument=IBM, Shares=1000, Price=25, gives TotalPrice=25000",
			args:           args{[]Bond{{"IBM", 1000, 25}}},
			wantTotalPrice: 25000,
		},
		{
			name:           "Instrument=IBM, Shares=1000, Price=25 and Instrument=GE, Shares=400, Price=100  gives TotalPrice=65000",
			args:           args{[]Bond{{"IBM", 1000, 25}, {"GE", 400, 100}}},
			wantTotalPrice: 65000,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := calculateTotalPrices(test.args.bonds); got != test.wantTotalPrice {
				t.Errorf("got=%v want=%v", got, test.wantTotalPrice)
			}
		})
	}
}

func TestConvertToUSD(t *testing.T) {
	type args struct {
		currency string
	}
	tests := []struct {
		name      string
		args      args
		want      string
		wantError bool
	}{
		{
			name: "10 CHF to USD at rate 1.5 gives 15 USD",
			args: args{"10 CHF"},
			want: "15 USD",
		},
		{
			name: "1 CHF to USD at rate 1.5 gives 1.5 USD",
			args: args{"1 CHF"},
			want: "1.5 USD",
		},
		{
			name: "80 INR to USD at rate 0.0125 gives 1 USD",
			args: args{"80 INR"},
			want: "1 USD",
		},
		{
			name:      "80 XYZ to USD when exchange rate are not available in table",
			args:      args{"80 XYZ"},
			want:      "",
			wantError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := convertToUSD(test.args.currency)
			if got != test.want {
				t.Errorf("got=%v want=%v", got, test.want)
			}
			if (err != nil) != test.wantError {
				t.Errorf("got error %s", err.Error())
			}
		})
	}
}
