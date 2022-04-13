package vendingmachine_test

import (
	"reflect"
	. "tdd/vendingmachine"
	"testing"
)

func TestGetChange(t *testing.T) {
	type args struct{
		totalPayable, cashPaid int
	}
	tests := []struct {
		name string
		args args
		wantChange []int
		wantErr bool
	}{
		{
			name: "getChange(1,1) should equal [] - an empty array",
			args: args{1,1},
			wantErr: false,
		},
		{
			name: "getChange(215,300) should return [50, 20, 10, 5]",
			args: args{215,300},
			wantChange: []int{50, 20, 10, 5},
			wantErr: false,
		},
		{
			name: "getChange(486, 600) should return [100, 10, 2, 2]",
			args: args{486, 600},
			wantChange: []int{100, 10, 2, 2},
			wantErr: false,
		},
		{
			name: "getChange(12, 400) should return [200, 100, 50, 20, 10, 5, 2, 1]",
			args: args{12, 400},
			wantChange: []int{200, 100, 50, 20, 10, 5, 2, 1},
			wantErr: false,
		},
		{
			name: "getChange(50,25) should return error",
			args: args{50, 25},
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got,err := GetChange(test.args.totalPayable, test.args.cashPaid)
			if !reflect.DeepEqual(got,test.wantChange) {
				t.Errorf("got=%v wantChange=%v", got, test.wantChange)
			}
			if (err!=nil)==!test.wantErr{
				t.Errorf("err=%v wantErr=%v", err, test.wantErr)
			}
		})
	}
}


