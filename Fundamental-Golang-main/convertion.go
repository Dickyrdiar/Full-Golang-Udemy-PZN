package main

import "fmt"

func main() {
	var nilai32 int32 = 129 // maksimum value 127 karena int 8 di akhir
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "via"
	var e byte = name[0]           //keluar byte
	var eString string = string(e) //konversi byte ke string

	fmt.Println(name)
	fmt.Println(eString)
}
