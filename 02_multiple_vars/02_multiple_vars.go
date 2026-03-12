package main

import "fmt"

func main() {
	var coffeName = "Espresso"
	var size = "Medium"
	price := 2.5

	fmt.Println("Medium Espresso price is $2.50")
	fmt.Println(size, coffeName, "price is $", price)
	fmt.Printf("%s %s price is $%.2f\n", size, coffeName, price)
}