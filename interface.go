package main

import "fmt"

type HasName interface {
	GeTName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GeTName())
}

type Person struct {
	Name string
}

func (person Person) GeTName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GeTName() string {
	return animal.Name
}

func main() {
	var sul Person
	sul.Name = "Sul"

	SayHello(sul)

	cat := Animal{
		Name: "push",
	}
	SayHello(cat)
}
