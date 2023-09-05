package main

import "fmt"

func modifyMap(m map[string]int) {
	m["change"] = 5
}

func main() {
	m := map[string]int{}
	modifyMap(m)
	fmt.Println(m)
}
