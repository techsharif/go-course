package main

import "fmt"

func createIncrementer(incrementBy int) func(int) {
	return func(x int) {
		fmt.Println(x + incrementBy)
	}
}

func main() {
	incrementByThree := createIncrementer(3)
	incrementByThree(7)
	incrementByThree(9)

	incrementByFive := createIncrementer(5)
	incrementByFive(2)

}
