package main

import "fmt"

func main() {

	// basic switch
	// num := 4

	// switch num {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println("Number not recognized")
	// }

	// Multiple Expressions in Case
	// day := "Sunday"

	// switch day {
	// case "Sun", "Mon", "Tue", "Wed", "Thu":
	// 	fmt.Println("It's a weekday")
	// case "Fri", "Sat":
	// 	fmt.Println("It's the weekend")
	// default:
	// 	fmt.Println("Invalid day")
	// }

	// Switch with No Expression
	// num := 0

	// switch {
	// case num > 0:
	// 	fmt.Println(num, "is a positive number")
	// case num < 0:
	// 	fmt.Println(num, "is a negative number")
	// default:
	// 	fmt.Println(num, "is zero")
	// }

	// Fallthrough Keyword
	num := 2
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Number not recognized")
	}
}
