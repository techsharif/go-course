package main

import "fmt"

func modifySlice(s *[]int) {
	fmt.Println(*s, &(*s)[0])
	*s = append(*s, 4)
	fmt.Println(*s, &(*s)[0])
}

func main() {
	s1 := &[]int{1, 2, 3}
	fmt.Println(*s1, &(*s1)[0])
	modifySlice(s1)
	fmt.Println(*s1, &(*s1)[0])
}
