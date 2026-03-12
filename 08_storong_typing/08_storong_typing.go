package main

import "fmt"

func main(){
	// price of one cup of coffee
	price := 4.50

	// cups sold in one day
	quantity := 15

	// total income
	total := price * float64(quantity)
	fmt.Printf("Total income during a day %.2f\n", total)
}