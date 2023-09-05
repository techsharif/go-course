package main

import "fmt"

func PrintType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("It's an integer")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Different type")
	}
}

func main() {
	PrintType(5)
	PrintType("Hello")
	PrintType(1.2)
}
