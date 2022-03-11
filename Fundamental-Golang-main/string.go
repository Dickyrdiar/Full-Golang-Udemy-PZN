package main

import "fmt"

func main() {
	fmt.Println("Octavia")
	fmt.Println("Octavia via")
	fmt.Println("Octavia via via")

	fmt.Println(len("via"))
	fmt.Println("octavia via"[0])
	fmt.Println("octavia via via"[1]) //keluar byte nya c = 99
}

//note
/* len untuk jumlah kata via=3
karakter di mulai dari 0 1 2 3 4 dst
*/
