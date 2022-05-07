package main

import "fmt"

func main() {
	var sayi1 = 5 // int tipinde degisken tanimlama
	var sayi2 int
	var sayi3 float32 //float tipinde değilken tanımlama
	var isim string   // metinsel veri tipinde veri tipi tanımalam
	var evli bool     // true false tipinde veri tanımlama

	var oylama float32 = 4.5 //float32 tipinde veri tanımala ve değer ekelem

	fmt.Println(sayi1)  // çıktısı => 5
	fmt.Println(sayi2)  // çıktısı => 0
	fmt.Println(sayi3)  // çıktısı => 0
	fmt.Println(isim)   // çıktısı ""
	fmt.Println(evli)   // çıktısı => false
	fmt.Println(oylama) // çıktısı => 4.5
}
