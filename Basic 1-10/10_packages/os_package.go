package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("test.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	file.WriteString("Hello World")

	fmt.Println("File created successfully")
}
