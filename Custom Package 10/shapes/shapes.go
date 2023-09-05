package main

import (
	"fmt"
	"shapes/rectangle"
)

func main() {
	r := rectangle.Rectangle{2, 2}
	area := r.area()
	fmt.Println(area)
}
