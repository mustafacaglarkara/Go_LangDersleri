package main

import (
	"fmt"
)

var isim, soyisim = "KARA", "Mustafa Çağlar"

func main() {
	isim, soyisim = soyisim, isim
	fmt.Println(isim, soyisim)
	fmt.Println(slmVer())
}

func slmVer() (msj string) {
	msj = "Merhabalar"
	return
}
