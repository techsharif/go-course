package main

import "fmt"

type Dog struct {
	Name  string
	Age   int
	Color string
}

func (d Dog) Bark() {
	fmt.Printf("%s says: woof!\n", d.Name)
}

func (d *Dog) ChangeName(name string) {
	d.Name = name
}

func main() {
	d := Dog{Name: "Rex", Age: 5, Color: "Brown"}
	d.ChangeName("Tix")
	d.Bark()
}
