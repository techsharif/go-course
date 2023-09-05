package main

import "fmt"

func main() {
	var x int = 2
	var p *int
	p = &x
	fmt.Println(*p)

	*p = 5

}
