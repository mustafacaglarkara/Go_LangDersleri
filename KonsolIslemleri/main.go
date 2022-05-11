package main

import "fmt"

var str1 = "Deneme Mesaj 1"
var str2 = "Mesaj 2"
var str3 = "Mesaj 30"
var aNumber = 38
var isTrue = true

func main() {

	/*
		Ders2 bölüm iki de detaylı var
	*/
	// int olara metnin kaç karakter olduğunu  belirtiyor.
	stringLenght, err := fmt.Println(str1, str2, str3)
	if err == nil {
		fmt.Println(stringLenght)
	}
	fmt.Printf("Parasal olarak değer : %.2f \n", float32(aNumber))

}
