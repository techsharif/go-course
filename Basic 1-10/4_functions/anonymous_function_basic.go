package main

import "fmt"

// func hello() {
// 	fmt.Println("Hello")
// }

func main() {

	result := func(x int, y int) int {
		return x + y
	}(5, 6)
	fmt.Println(result)

	// func (){
	// 	fmt.Println("Hello")
	// }()

	// result := sum(5, 6)
	// fmt.Println(result)

	// anotherHello := hello
	// anotherHello()
}
