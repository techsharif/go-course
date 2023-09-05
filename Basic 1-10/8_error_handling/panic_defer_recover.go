package main

import "fmt"

func deferFunc() {
	fmt.Println("defer function")
	if r := recover(); r != nil {
		fmt.Println("recover from", r)
	}
}

func main() {
	defer deferFunc()

	fmt.Println("Before panic")
	panic("something bad happend")
	fmt.Println("After panic")
}
