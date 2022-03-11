package main

import "fmt"

func main() {
	var name string //satu kali bikin untuk satu program

	name = "Octavia via"
	fmt.Println(name)

	name = "Octavia via via"
	fmt.Println(name)

	var friendName = "via" //tidak perlu menyebutkan tipe data /string
	fmt.Println(friendName)

	var age = 22
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	var (
		firstName = "octavia"
		lastName  = "via"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
