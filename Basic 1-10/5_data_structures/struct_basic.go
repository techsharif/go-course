package main

import "fmt"

func main() {

	// creating a struct
	type Person struct {
		Name string
		Age  int
	}

	// var p Person
	// fmt.Println(p)

	// Assigning values
	var p Person = Person{"Abul", 25}
	fmt.Println(p)

}
