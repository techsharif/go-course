package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Println(i)
	}
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c\n", i)
	}
}

func main() {
	go printNumbers()
	go printLetters()

	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Main function finished")
}
