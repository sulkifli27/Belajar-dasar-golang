package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Sul"
	names[1] = "Kifli"
	names[2] = "Sulkifli"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		87,
		75,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))

}
