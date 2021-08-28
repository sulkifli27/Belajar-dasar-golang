package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Sulkifli" , "Sul"))
	fmt.Println(strings.Contains("Sulkifli" , "Jul"))

	fmt.Println(strings.Split("Sul kifli sul", " "))

	fmt.Println(strings.ToLower("Sul kifli sul"))
	fmt.Println(strings.ToUpper("Sul kifli sul"))
	fmt.Println(strings.ToTitle("sul kifli sul"))

	fmt.Println(strings.Trim("      sul kifli sul      ", " "))
	fmt.Println(strings.ReplaceAll("sul ijul sul", "sul", "Kifli"))
}