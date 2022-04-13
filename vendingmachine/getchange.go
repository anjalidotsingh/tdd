package vendingmachine

import (
	"fmt"
	"math"
)

//coins are arranged in decending order
var coins = []int{200, 100, 50, 20, 10, 5, 2, 1}

func GetChange(totalPayable, cashPaid int) ([]int, error) {
	var change []int
	returnAmount := cashPaid - totalPayable
	if returnAmount == 0 {
		return change, nil
	}
	if returnAmount < 0 {
		return change, fmt.Errorf("amount paid is less than price of item")
	}
	for _, coin := range coins {
		if coin > returnAmount {
			continue
		}
		noOfCoin:=int(math.Floor(float64(returnAmount)/float64(coin)))
		for i:=1;i<=noOfCoin;i++{
			change = append(change, coin)
		}
		returnAmount = returnAmount-(coin*noOfCoin)
		if returnAmount == 0 {
			break
		}
	}
	return change, nil
}
