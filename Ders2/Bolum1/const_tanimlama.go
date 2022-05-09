package main

import "fmt"

/*
Diğer dillerde sabit sanımlarken genelde büyük harf ile
başlarlar.
Go da büyük harf ile başmamasına gerek yoktur.
*/
const egitmen = "Mustafa Çağlar KARA"

const (
	pazartesi = iota // otomatik olarak 0 dan başlayıp diğerlerine de değer verir
	sali
	carsamba
	persembe
	cuma
	cumartesi
	pazar
)

func main() {
	fmt.Println(egitmen)
	fmt.Println(pazartesi, sali)
}
