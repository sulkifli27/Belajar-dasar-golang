package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan Ke", counter)
	// 	counter++
	// }

	for counter := 1; counter < 10; counter++ {
		fmt.Println("Perulangan Ke", counter)
	}

	slice := []string{"Sulkifli", "Sul", "Kifli", "Ijul", "Ijul sul "}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for _, value := range slice {
		// fmt.Println("Index", i, "=", value)
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Sulkifli"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
