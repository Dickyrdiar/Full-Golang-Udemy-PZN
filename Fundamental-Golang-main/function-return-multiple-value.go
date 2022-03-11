package main

import "fmt"

func getFullName() (string, string, string) {
	return "Via", "Octavia", "vea"
}

func main() {
	firstName, _, _ := getFullName()
	fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(lastName)
}

// _ ; tidak dipedulikan
