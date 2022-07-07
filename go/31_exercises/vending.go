package main

import (
	"fmt"
)

var coins = []int{1, 5, 10, 25, 50, 100}

func main() {
	var amount float32 = 5.0
	var price float32 = 1.00
	change := getChange(amount, price)
	fmt.Printf("Change for %.2f when paying for %.2f item: %v\n", amount, price, change)

	amount = 5.0
	price = 0.99
	change = getChange(amount, price)
	fmt.Printf("Change for %.2f when paying for %.2f item: %v\n", amount, price, change)

	amount = 3.14
	price = 0.99
	change = getChange(amount, price)
	fmt.Printf("Change for %.2f when paying for %.2f item: %v\n", amount, price, change)

	amount = 3.14
	price = 2.83
	change = getChange(amount, price)
	fmt.Printf("Change for %.2f when paying for %.2f item: %v\n", amount, price, change)
}

func getChange(amount float32, price float32) []int {
	return getChangeInt(int(amount * 100), int(price * 100))
}

func getChangeInt(amount int, price int) []int {
	slice := make([]int, 6)
	change := amount - price
	for i := len(coins) - 1; i >= 0; i-- {
		for change >= coins[i] {
			slice[i]++
			change -= coins[i]
		}
	}
	return slice
}


