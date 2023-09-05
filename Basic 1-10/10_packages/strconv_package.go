package main

import (
	"fmt"
	"strconv"
)

func main() {
	bStr := "f"
	b, err := strconv.ParseBool(bStr)
	if err != nil {
		fmt.Println("error found", err)
	}

	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
