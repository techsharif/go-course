package main

import "fmt"

func changeTo5(numPointer *int) {
	*numPointer = 5
}

func main() {
	// var num int = 10
	// fmt.Println(num, &num)
	// var numPointer *int = &num
	// fmt.Println(numPointer, *numPointer)

	// // var num2 int = num
	// // fmt.Println(num2, &num2)

	// *numPointer = 20
	// fmt.Println(num)

	num := 10
	changeTo5(&num)
	fmt.Println(num)
}
