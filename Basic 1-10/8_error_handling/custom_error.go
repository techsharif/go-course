package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Msg  string
	Code int
}

func (err CustomError) Error() string {
	return fmt.Sprintf("Error with code %d: %s", err.Code, err.Msg)
}

func SuperDivide(a, b int) (int, error) {
	if b == 0 {
		// return 0, CustomError{Msg: "divide by zero", Code: 1}
		return 0, errors.New("Divide by zero")
	}

	if a%b != 0 {
		return 0, CustomError{Msg: "not divisible", Code: 2}
	}

	return a / b, nil
}

func main() {
	res, err := SuperDivide(5, 2)
	if err != nil {
		var cErr CustomError
		if errors.As(err, &cErr) {
			fmt.Println("Custom error occurred")
			fmt.Println("Code", cErr.Code, "Message", cErr.Msg)
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(res)
}
