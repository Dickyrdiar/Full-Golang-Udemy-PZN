package main

import "fmt"

func main() {

	var ujian = 80
	var absensi = 79 // false; 80 = lulus /true

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println(lulus)

	fmt.Println(ujian >= 80 && absensi >= 80)

}

/*
operasi bool cuma ada 3
&& and -> bernilai true jika keduanya true
|| or -> minimal salah satu true maka bernilai true
! negasi -> kebalikan

*/
