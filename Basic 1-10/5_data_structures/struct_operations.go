package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"abul", 25}
	fmt.Println(p)

	// Accessing fields
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)

	// Modifying fields
	p.Name = "Mr. Abul"
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)

}
