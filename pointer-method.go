package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	sul := Man{"Sul"}
	sul.Married()
	fmt.Println(sul.Name)
}
