package main

import "fmt"

func sum(nums ...int) {
	result := 0
	for _, num := range nums {
		result = result + num
	}
	fmt.Println(result)
}

func main() {
	sum(5, 6)    // 11
	sum(1, 2, 3) // 6

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
