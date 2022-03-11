package main

import "fmt"

func main() {

	name := "octavia"

	switch name {
	case "via":
		fmt.Println("Hello via")
		fmt.Println("Hello via")
	case "octavia":
		fmt.Println("Hello octa")
		fmt.Println("Hello octa")
	default:
		fmt.Println("Kenalan Donk")
		fmt.Println("Kenalan Donk")
	}

	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Nama terlalu panjang")
	//case false:
	//	fmt.Println("Nama sudah benar")
	//}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}

//swicth pengecekan sederhana
