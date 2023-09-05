package main

import "fmt"

func main() {

	// declaration
	// var arrayName [5]int
	// fmt.Println(arrayName)
	// fmt.Println(len(arrayName))
	// fmt.Printf("%T\n", arrayName)

	// Initialization

	// var arrayName [5]int = [5]int{10, 20, 30, 40, 50}
	// var arrayName = [5]int{10, 20, 30, 40, 50}
	// arrayName := [5]int{10, 20, 30, 40, 50}

	arrayName := [...]int{10, 20, 30, 40, 50}

	fmt.Println(arrayName)
	fmt.Println(len(arrayName))
	fmt.Printf("%T\n", arrayName)

}
