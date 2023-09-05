package main

import (
	"fmt"
	"time"
)

func getNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go getNumbers(ch)

	for num := range ch {
		fmt.Println(num)
	}
}
