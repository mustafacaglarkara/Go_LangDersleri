package main

import (
	"fmt"
)

var denemeBosDizi [3]int //boş olarak dizi tanımla
var sinavlar = [4]int{
	5,
	3,
	5,
	2,
}

func ortalamaHesapla(exam []int) {
	toplam := 0
	ortalama := 0.0
	for _, deger := range exam {
		toplam += deger
	}
	ortalama = float64(toplam / len(sinavlar))
	fmt.Println(toplam, ortalama)
}

func main() {
	// arrat to slice yapılması gerekli
	var gecis = sinavlar[:]
	ortalamaHesapla(gecis)
}
