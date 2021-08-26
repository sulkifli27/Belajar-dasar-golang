package main

import "fmt"

func main() {
	name := "Sul"

	switch name {
	case "Sulkifli":
		fmt.Println("Hello Sulkifli")
		fmt.Println("Hello Sulkifli")
	case "Sul":
		fmt.Println("Hello Sul")
		fmt.Println("Hello Sul")
	default:
		fmt.Println("Kenalan Dong")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayang panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
