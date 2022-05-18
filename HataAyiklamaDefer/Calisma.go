package main

import "fmt"

func main() {
	var yas = 0
	fmt.Println("Lütfen bir sayı giriniz : ")
	fmt.Scanln(&yas)

	YasKontrol(&yas)
	YasArttir(&yas)
	fmt.Println(yas)
	YasKontrol(&yas)
	BolmeKontrol(&yas)
}

func BolmeKontrol(i *int) {
	hata := recover()
	if hata != nil {
		fmt.Println("Hata Var")
	} else {
		fmt.Println("Hata Yok")
	}
}

func YasArttir(i *int) {
	*i += 5
}

func YasKontrol(sayi *int) {
	if *sayi > 18 {
		fmt.Println("Yaşınız 18 den büyük")
	} else {
		fmt.Println("Yaşınız 18 den küçük")
	}
}
