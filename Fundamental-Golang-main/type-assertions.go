package main

import "fmt"

func random() interface{} {
	return "via"
}

func main() {
	var result interface{} = random()
	//var resultString string = result.(string)
	//fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}

/*
type asssertion mengubah jenis data
jika salah menggunakan type ass -> panic-> program mati
*/
