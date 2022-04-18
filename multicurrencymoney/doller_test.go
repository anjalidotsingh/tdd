package multicurrencymoney_test

import (
	. "tdd/multicurrencymoney"
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
-refractor
-removing magic constants
*/

func TestCalculateTotalPrice(t *testing.T) {
	type args struct {
		bonds []Bond
	}
	tests := []struct {
		name           string
		args           args
		wantTotalPrice string
		wantError      bool
	}{
		{
			name:           "Instrument=IBM, Shares=1000, Price=25 USD, gives TotalPrice=25000",
			args:           args{[]Bond{{"IBM", 1000, "25 USD"}}},
			wantTotalPrice: "25000 USD",
		},
		{
			name:           "Instrument=IBM, Shares=1000, Price=25 USD and Instrument=GE, Shares=400, Price=100 USD gives TotalPrice=65000",
			args:           args{[]Bond{{"IBM", 1000, "25 USD"}, {"GE", 400, "100 USD"}}},
			wantTotalPrice: "65000 USD",
		},
		{
			name:           "Instrument=IBM, Shares=1000, Price=25 USD and Instrument=Novartis, Shares=400, Price=150 CHF `at rate CHF:USD::1.5:1 gives TotalPrice=115000",
			args:           args{[]Bond{{"IBM", 1000, "25 USD"}, {"GE", 400, "150 CHF"}}},
			wantTotalPrice: "115000 USD",
		},
		{
			name:      "Instrument=SBI, Shares=10, Price=25 XYZ exchange rate are not available gives error",
			args:      args{[]Bond{{"SBI", 10, "25 XYZ"}}},
			wantError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := CalculateTotalPrices(test.args.bonds)
			if got != test.wantTotalPrice {
				t.Errorf("got=%v want=%v", got, test.wantTotalPrice)
			}
			if (err != nil) != test.wantError {
				t.Errorf("got error %s", err.Error())
			}

		})
	}
}

func TestConvertToUSD(t *testing.T) {
	type args struct {
		denomination float64
		unit         string
	}
	tests := []struct {
		name      string
		args      args
		want      float64
		wantError bool
	}{
		{
			name: "10 CHF to USD at rate CHF:USD::1.5:1 gives value as 15 and no error",
			args: args{10, "CHF"},
			want: 15,
		},
		{
			name: "1 CHF to USD at rate CHF:USD::1.5:1 gives value as 1.5 and no error",
			args: args{1, "CHF"},
			want: 1.5,
		},
		{
			name: "80 INR to USD at rate INR:USD::0.0125:1 gives value as 1 and no error",
			args: args{80, "INR"},
			want: 1,
		},
		{
			name:      "80 XYZ to USD when exchange rate are not available",
			args:      args{80, "XYZ"},
			want:      0,
			wantError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := ConvertToUSD(test.args.denomination, test.args.unit)
			if got != test.want {
				t.Errorf("got=%v want=%v", got, test.want)
			}
			if (err != nil) != test.wantError {
				t.Errorf("got error %s", err.Error())
			}
		})
	}
}
