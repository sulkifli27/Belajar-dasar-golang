package main

import "fmt"

func main() {
	type NoKtp string
	type Married bool

	var noKtpSul NoKtp = "8232193928932939"
	var marriedStatus Married = true
	fmt.Println(noKtpSul)
	fmt.Println(marriedStatus)

}
