package main

import (
	"fmt"
	"time"
)

func server1(ch1 chan string) {
	time.Sleep(4 * time.Second)
	ch1 <- "from server 1"
}

func server2(ch2 chan string) {
	time.Sleep(2 * time.Second)
	ch2 <- "from server 2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)

	select {
	case s1 := <-output1:
		fmt.Println(s1)

	case s2 := <-output2:
		fmt.Println(s2)
	default:
		fmt.Println("No data found")
	}
}
