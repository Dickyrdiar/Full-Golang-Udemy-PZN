package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "octa"
	sayHelloTo(firstName, "via")
	sayHelloTo("otta", "vea")
}
