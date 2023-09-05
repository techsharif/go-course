package main

import "fmt"

func main() {

	// creating slice
	// var mySlice []int
	// fmt.Println(mySlice)
	// fmt.Println(len(mySlice))
	// fmt.Printf("%T\n", mySlice)

	// slice from array
	// arr := [5]int{10, 20, 30, 40, 50}
	// mySlice := arr[1:4]
	// fmt.Println(mySlice)
	// fmt.Println(len(mySlice))
	// fmt.Printf("%T\n", mySlice)

	// slice initialization
	// var mySlice []int = []int{100, 200, 300}
	// mySlice := []int{100, 200, 300}
	// fmt.Println(mySlice)
	// fmt.Println(len(mySlice))
	// fmt.Printf("%T\n", mySlice)

	// using make
	mySlice := make([]int, 3)

	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Printf("%T\n", mySlice)

}
