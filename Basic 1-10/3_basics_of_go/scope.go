package main

import "fmt"

// global
// local

var globalVar = "global value"

func main() {
	var localVar = "local value"

	fmt.Println(globalVar)
	fmt.Println(localVar)
}

func anotherFunc() {
	fmt.Println(globalVar)
	fmt.Println(localVar)
}
