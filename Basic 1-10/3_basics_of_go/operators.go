package main

import "fmt"

func main() {
	// Arithmetic operators
	// a := 6
	// b := 4

	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)
	// fmt.Println(a % b)

	// a++
	// b--
	// fmt.Println(a, b)

	// Relational operators
	// a := 5
	// b := 6
	// c := 5

	// ==
	// fmt.Println(a == c)
	// fmt.Println(a == b)

	// !=
	// fmt.Println(a != b)
	// fmt.Println(a != c)

	// >
	// fmt.Println(a > b)
	// fmt.Println(b > c)
	// fmt.Println(a > c)

	// <
	// fmt.Println(a < b)
	// fmt.Println(b < c)
	// fmt.Println(a < c)

	// >=
	// fmt.Println(a >= b)
	// fmt.Println(b >= c)
	// fmt.Println(a >= c)

	// <=
	// fmt.Println(a <= b)
	// fmt.Println(b <= c)
	// fmt.Println(a <= c)

	// logical operator
	// t := true  // 5 == 5
	// f := false // 5 > 6
	// t1 := true
	// f1 := false

	// // && logical and
	// fmt.Println(t && t1)
	// fmt.Println(t && f)
	// fmt.Println(f && t)
	// fmt.Println(f && f1)

	// || logical or
	// fmt.Println(t || t1)
	// fmt.Println(t || f)
	// fmt.Println(f || t)
	// fmt.Println(f || f1)

	// ! logical not
	// fmt.Println(!t)
	// fmt.Println(!f)

	// n -> 3 and 5 divisible or not
	// n := 6
	// div3 := n%3 == 0
	// div5 := n%5 == 0
	// result := div3 && div5
	// fmt.Println(result)

	// bitwise operator
	// a := 5 // 0101
	// b := 3 // 0011

	// fmt.Println(a & b) // & 0001 = 1
	// fmt.Println(a | b) // | 0111 = 7
	// fmt.Println(a ^ b) // ^ 0110 = 6

	// fmt.Println(^a) // ^ 1010 = -6

	// assignment operator
	a := 5

	a = 6
	fmt.Println(a) // 6

	a += 5         // a = a+5
	fmt.Println(a) // 11

	a -= 2         // a = a-2
	fmt.Println(a) // 9

	a *= 3         // a = a*3
	fmt.Println(a) // 27

	a /= 7         // a = a/7
	fmt.Println(a) // 3

	a %= 7         // a = a%7
	fmt.Println(a) // 3
}
