package main

import "fmt"

func main() {

	// basic loop
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// For with just condition
	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// Infinite loop
	// i := 1
	// for {
	// 	fmt.Println(i)
	// 	i++
	// 	if i > 5 {
	// 		break
	// 	}
	// }

	// For loop with range
	numbers := []int{1, 2, 3, 4, 5}
	for i, num := range numbers {
		fmt.Printf("Index %d, Number: %d\n", i, num)
	}
}
