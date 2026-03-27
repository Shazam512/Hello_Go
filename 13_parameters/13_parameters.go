package main

import "fmt"

func greet(name string) {
	fmt.Println("Welcome to Go programming!", name)
}

func main() {	
	greet("This is my first test")
	greet("This is my two test")
}
