package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	result := 10 / value
	fmt.Println("Result", result)
	fmt.Println("Run Application")
}

func main() {
	runApplication(10)
}
