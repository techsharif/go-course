package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)

	specificTime := time.Date(2023, 07, 04, 0, 0, 0, 0, time.UTC)
	fmt.Println(specificTime)
}
