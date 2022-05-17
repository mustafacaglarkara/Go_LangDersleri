package main

import "fmt"

func main() {
	fmt.Println(Hesapla(10))
}
func A() {
	fmt.Println("A Fonksiyonu Çalıştı")
}
func B() {
	fmt.Println("B Fonksiyonu Çalıştı")
}
func Hesapla(sayi int) string {
	defer B()
	if sayi%2 == 0 {
		fmt.Println("Çift Sayı Çalıştı")
		return "Çift Sayı"
	}
	if sayi%2 != 0 {
		fmt.Println("Tek Sayı Çalıştı")
		return "Tek Sayı"
	}
	return "Bitti"
}
