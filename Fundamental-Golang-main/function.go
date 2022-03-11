package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World 1")
	fmt.Println("Hello World 2")
	fmt.Println("Hello World 3")
}

func main() {
	for i := 0; i < 10; i++ {
		sayHello()
	}
}
