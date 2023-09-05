package main

import "fmt"

func main() {

	arr := [5]int{10, 20, 30, 40, 50}

	// Accessing Elements
	fmt.Println(arr[2])

	// Modifying Elements
	arr[2] = 60
	fmt.Println(arr)

	// Iterating over an Array
	for i, v := range arr {
		fmt.Printf("index: %d, value %d\n", i, v)
	}
}
