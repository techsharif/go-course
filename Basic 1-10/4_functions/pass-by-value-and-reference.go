package main

import "fmt"

func passByValue(x int) {
	x = 10
}

func passByReference(y *int) {
	*y = 20
}

func main() {
	// pass by value
	// value := 5
	// passByValue(value)
	// fmt.Println(value)

	// pass by reference
	value := 5
	passByReference(&value)
	fmt.Println(value)
}
