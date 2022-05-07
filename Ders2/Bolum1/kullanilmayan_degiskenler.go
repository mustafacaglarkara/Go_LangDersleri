package main

import "fmt"

func main() {
	var isim = "Mustafa" //kullanılmadığından hata verecektir.
	_ = isim             // tanımlanır ise hata vermez.
	fmt.Println("İsim adlı değişlken kullanılmadığından hata verir.")
}
