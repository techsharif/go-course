package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Print() {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

func (p *Person) CelebrateBirthday() {
	p.Age++
	fmt.Printf("Happy birthday %s, you are now %d years old\n", p.Name, p.Age)
}

func main() {
	p := Person{"abul", 25}
	p.Print()
	p.CelebrateBirthday()
	p.Print()
	// p.Print()

}
