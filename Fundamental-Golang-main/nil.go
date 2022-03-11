package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var person map[string]string = NewMap("via")

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}
}

//objek yang belum di iniliasisasi maka akan menggunakan defauld value
// nil => hanya untuk interface, function, map, slice, pointer,channel
