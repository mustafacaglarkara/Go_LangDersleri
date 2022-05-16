package main

import (
	"fmt"
)

func main() {
	var sayi1, sayi2, sayi3 = 10, 55, 12588

	var enBuyukSayi = sayi1

	if sayi2 > enBuyukSayi {
		enBuyukSayi = sayi2
	}
	if sayi3 > enBuyukSayi {
		enBuyukSayi = sayi3
	}
	fmt.Println(sayi1, sayi2, sayi3)
	fmt.Println("En büyük Sayı : ", enBuyukSayi)
}
