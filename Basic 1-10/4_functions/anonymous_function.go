package main

import "fmt"

// func hello() {
// 	fmt.Println("Hello")
// }

func execute(x int, y int, operation func(x int, y int) int) {
	result := operation(5, 6)
	fmt.Println(result)
}

func main() {

	sum := func(x int, y int) int {
		return x + y
	}
	execute(5, 6, sum)
	execute(5, 6, func(a int, b int) int {
		return a * b
	})

	// result := sum(5, 6)
	// fmt.Println(result)

	// anotherHello := hello
	// anotherHello()
}
