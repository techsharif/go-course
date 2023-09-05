package main

import "fmt"

func modifyArray(a *[5]int) {
	fmt.Println((*a)[0], &(*a)[0])
	(*a)[0] = 5
}

func main() {
	arr := &[5]int{1, 2, 3, 4, 5}
	fmt.Println((*arr)[0], &(*arr)[0])
	modifyArray(arr)
	fmt.Println(arr)
}
