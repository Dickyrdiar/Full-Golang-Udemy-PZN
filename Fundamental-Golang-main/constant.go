package main

import "fmt"

func main() {
	const ( //wajib langsung deklarasi value nya
		firstName string = "octavia" //tidak boleh pakai :=
		lastName         = "via"
		value            = 1000
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
}
