package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	res, err := divide(4, 0)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	fmt.Println(res)
}
