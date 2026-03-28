package main

import "fmt"

func calcPoints(amountSpent float64) int {
	points := int(amountSpent * 2)
	return points
}

func main(){
	var newlyPoints int = calcPoints(9.50)
	fmt.Println("Earned points today:", newlyPoints)
}