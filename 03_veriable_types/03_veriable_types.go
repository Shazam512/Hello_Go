package main

import "fmt"

func main(){
	name := "Americano"
	prise := 2.99
	ready := true
	count := 5

	fmt.Printf("Tupe on nema: %T\n", name)
	fmt.Printf("Tupe on prise: %.2f\n", prise)
	fmt.Printf("Tupe on ready: %T\n", ready)
	fmt.Printf("Tupe on count: %T\n", count)
}