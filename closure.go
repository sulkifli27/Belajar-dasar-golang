package main

import "fmt"

func main() {
	name := "Sul"
	counter := 0

	increment := func() {
		name := "Ijul"
		fmt.Println("Increment")
		fmt.Println(name)
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)

}
