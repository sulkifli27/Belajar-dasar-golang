package main

import "fmt"

func random() interface{} {
	return 100
}

func main() {
	var result interface{} = random()
	// var resultString = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("value", value, "is string")
	case int:
		fmt.Println("value", value, "is int")
	default:
		fmt.Println("unknow type")
	}

}
