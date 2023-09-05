package main

import "fmt"

func main() {
	const a int = 5
	fmt.Println(a)

	// iota

	const (
		First = iota
		Second
		Third
	)

	fmt.Println(First, Second, Third)
}
