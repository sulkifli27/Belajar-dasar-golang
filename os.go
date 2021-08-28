package main

import (
	"fmt"
	"os"
)

func main() {
	argss := os.Args
	fmt.Println("Argumen")
	fmt.Println(argss)

	name, err := os.Hostname()

	if err ==nil {
		fmt.Println("Hostname", name)
	}else{
		fmt.Println("Error", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)

}