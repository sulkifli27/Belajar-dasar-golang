package main

import "fmt"

func main() {
	var name = "Ijul"

	if name == "Sul" {
		fmt.Println("Hello Sul")
	} else if name == "Kifli" {
		fmt.Println("Hello Kifli")
	} else if name == "Ijul" {
		fmt.Println("Hello Ijul")
	} else {
		fmt.Println("Hi, kenalan yuk ?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
