package main

import (
	"fmt"
	"time"
)

func Multiply(a int, b int, ch chan int) {
	time.Sleep(250 * time.Millisecond)
	ch <- a * b
}

func main() {
	ch := make(chan int)

	go Multiply(2, 3, ch)

	res := <-ch
	fmt.Println(res)

	close(ch)

	res = <-ch
	fmt.Println(res)
}
