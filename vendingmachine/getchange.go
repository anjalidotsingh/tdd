package vendingmachine

import (
	"fmt"
	"math"
)

//coins are arranged in decending order
var coins = []int{200, 100, 50, 20, 10, 5, 2, 1}

func GetChange(totalPayable, cashPaid int) ([]int, error) {
	var returnCoins []int
	difference := cashPaid - totalPayable
	if difference == 0 {
		return returnCoins, nil
	}
	if difference < 0 {
		return returnCoins, fmt.Errorf("amount paid is less than price of item")
	}
	returnCoins=getReturnCoins(difference)
	return returnCoins, nil
}

func appendCoinToReturnAmount(returnCoins []int,coin,noOfCoin int)[]int{
	for i:=1;i<=noOfCoin;i++{
		returnCoins = append(returnCoins, coin)
	}
	return returnCoins
}

func getReturnCoins(difference int)[]int{
	var returnCoins []int
	for _, coin := range coins {
		if coin > difference {
			continue
		}
		noOfCoin:=int(math.Floor(float64(difference)/float64(coin)))
		returnCoins=appendCoinToReturnAmount(returnCoins,coin,noOfCoin)
		difference= difference-(coin*noOfCoin)
		if difference == 0 {
			break
		}
	}
	return returnCoins
}