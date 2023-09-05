package main

import "fmt"

func add(x int, y int) int {
	fmt.Println("function add called")
	return x + y
}

func main() {
	result := add(5, 6)
	fmt.Println(result)
}
