package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "Sul"
	sayHelloTo(firstName, "Kifli")
	sayHelloTo("Ijul", "Kifli")
}
