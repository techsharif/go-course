package main

import "fmt"

type Walker interface {
	Walk()
}

type Dog struct {
	Name  string
	Age   int
	Color string
}

func (d Dog) Walk() {
	fmt.Println("Dog with name", d.Name, "is walking")
}

type Cat struct {
	Name  string
	Age   int
	Color string
}

func (c Cat) Walk() {
	fmt.Println("Cat with name", c.Name, "is walking")
}

func MakeWalk(w Walker) {
	w.Walk()
}

func main() {
	dog := Dog{Name: "Rex", Age: 5, Color: "Brown"}
	MakeWalk(dog)

	cat := Cat{Name: "Whiskers", Age: 3, Color: "Black"}
	MakeWalk(cat)
}
