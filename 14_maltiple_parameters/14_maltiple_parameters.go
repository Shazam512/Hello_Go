package main

import "fmt"

func getDrinkInfo(custumerName string, drinkName string, getPrice float64) {
	fmt.Printf("%s's favorite drink is %s and it`s price is $%.2f\n", custumerName, drinkName, getPrice)
}

func main() {
	getDrinkInfo("Maxim", "Latte", 3.50)
	getDrinkInfo("Alex", "Capuccino", 4.00)
}