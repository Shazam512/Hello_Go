package main

import "fmt"

func main() {
	var coffeeType string = "Latte"
	var quantity int = 3
	var unitPrice float64 = 4.25

	fmt.Printf("Order: %d %s priced $%.2f each\n", quantity, coffeeType, unitPrice)

	var (
		customerName string = "Bogdan"
		tableNumber  int    = 5
		isReadyToPay bool   = false
	)

	fmt.Printf("Customer: %s, table %d, ready to pay: %t\n", customerName, tableNumber, isReadyToPay)

	const (
		SizeSmall  = "S"
		SizeMedium = "M"
		SizeLarge  = "L"
	)
}