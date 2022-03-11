package main

import "fmt"

func main() {
	var name = "raka"

	if name == "arnanda" {
		fmt.Println("Hello nanda")
	} else if name == "raka" {
		fmt.Println("Hello raka")
	} else if name == "octa" {
		fmt.Println("Hello octa")
	} else {
		fmt.Println("Hi, kenalan donk")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
