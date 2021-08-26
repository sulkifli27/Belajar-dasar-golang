package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayhuuu() {
	fmt.Println("Huuuuu from", a.Name)
}

func main() {
	var sul Customer
	sul.Name = "Sulkifli"
	sul.Address = "Indonesia"
	sul.Age = 23

	sul.sayHi("panjul")
	sul.sayhuuu()

	// fmt.Println(sul.Name)
	// fmt.Println(sul.Address)
	// fmt.Println(sul.Age)

	// ijul := Customer{
	// 	Name:    "Ijul",
	// 	Address: "Sinjai",
	// 	Age:     23,
	// }

	// fmt.Println(ijul)

	// kifli := Customer{"Kifli", "Sinjai", 23}
	// fmt.Println(kifli)

}
