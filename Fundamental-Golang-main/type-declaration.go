package main

import "fmt"

func main() {
	type NoKTP string //NoKTP = string
	type Married bool //Married = bool

	var noKtpvia NoKTP = "18741982741897419874"
	var marriedStatus Married = false
	fmt.Println(noKtpvia)
	fmt.Println(marriedStatus)
}

//mirip alias tipe data
