package multicurrencymoney

import (
	"testing"
)

/*
to-do list
1-add calculate price method 
2-add exchange rate
3-add support for multiple currency using exchange rate
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
			args:           args{[]Bond{{"IBM", 1000, 25},{"GE",400,100}}},
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
