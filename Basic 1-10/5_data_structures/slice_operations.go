package main

import "fmt"

func main() {

	// mySlice := []int{10, 20, 30, 40, 50}

	// accessing slice
	// fmt.Println(mySlice[2])

	// modify slice
	// mySlice[2] = 60
	// fmt.Println(mySlice[2])

	// appending elements
	// mySlice = append(mySlice, 60, 70, 80)
	// fmt.Println(mySlice)

	// Slice of Slice
	// newSlice := mySlice[2:7]
	// fmt.Println(newSlice)

	// iterate slice
	// for i, v := range mySlice {
	// 	fmt.Printf("Index %d, value %d\n", i, v)
	// }

	// copy a slice using equal
	// copyMySlice := mySlice
	// fmt.Println(mySlice)
	// fmt.Println(copyMySlice)
	// mySlice[2] = 60
	// fmt.Println(mySlice)
	// fmt.Println(copyMySlice)

	// copy slice
	mySlice := []int{10, 20, 30, 40, 50}
	copyMySlice := make([]int, len(mySlice))
	copy(copyMySlice, mySlice)
	fmt.Println(mySlice)
	fmt.Println(copyMySlice)
	mySlice[2] = 60
	fmt.Println(mySlice)
	fmt.Println(copyMySlice)

}
