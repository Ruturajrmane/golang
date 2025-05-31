package main

import (
	"fmt"
)

type Currency string

const (
	INR Currency = "INR"
	USD Currency = "USD"
	CAD Currency = "CAD"
)

var supportedCurrency = []Currency{INR, USD, CAD}

func main() {
	var money string
	fmt.Print("Enter the currency : ")
	fmt.Scanln(&money)

	moneyType := Currency(money)

	for _, curr := range supportedCurrency {
		if curr == moneyType {
			fmt.Println("This currency is acceptable")
			break
		} else {
			fmt.Printf("%s currency can't be accepted", money)
			return
		}
	}

	switch moneyType {
	case INR:
		fmt.Println("This currency is Indian Ruppes")
	case USD:
		fmt.Println("This currecny is US Dollar")
	case CAD:
		fmt.Println("This currency is Canadian Dollar")
	default:
		fmt.Println("This is Unkonown")
	}
}
