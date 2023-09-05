package main

import "fmt"

func main() {

	// creating maps
	// var m map[string]int
	// fmt.Println(m)

	// initialization
	var m map[string]int = map[string]int{
		"apple":  5,
		"banana": 10,
		"cherry": 15,
	}
	fmt.Println(m)

	// accessing elements
	// fmt.Println(m["banana"])

	// modifing elements
	// m["banana"] = 20
	// fmt.Println(m["banana"])

	// add new key
	// m["orange"] = 30
	// fmt.Println(m)

	// iterating elements
	// for k, v := range m {
	// 	fmt.Printf("Key %s, Value %d\n", k, v)
	// }

	// deleting elements
	delete(m, "banana")
	fmt.Println(m)
}
