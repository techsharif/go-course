package main

import "fmt"

func main() {

	// Short Statement with If
	// num := 10

	// if num > 0 {
	// 	fmt.Println(num, "is positive")
	// } else {
	// 	fmt.Println(num, "is not positive")
	// }

	// More than two conditions (else if)
	// num := 10
	// if num > 0 {
	// 	fmt.Println(num, "is positive")
	// } else if num == 0 {
	// 	fmt.Println(num, "is zero")
	// } else {
	// 	fmt.Println(num, "is negative")
	// }

	// Nested If
	// num := 0

	// if num == 0 {
	// 	fmt.Println(num, "is zero")
	// } else {
	// 	if num > 0 {
	// 		fmt.Println(num, "is positive")
	// 	} else {
	// 		fmt.Println(num, "is negative")
	// 	}
	// }

	// If with a Function
	num := -10
	if isPositive(num) {
		fmt.Println(num, "is positive")
	} else {
		fmt.Println(num, "is not positive")
	}

}

func isPositive(num int) bool {
	return num > 0
}
