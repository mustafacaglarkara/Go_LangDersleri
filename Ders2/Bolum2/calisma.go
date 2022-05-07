package main

import (
	"fmt"
	"strconv"
)

func main() {

	var input string
	fmt.Println("Lütfen Yaşınızı Giriniz")
	fmt.Scanf("%s", &input)
	yas, err := strconv.Atoi(input)
	age, err := strconv.ParseInt(input, 10, 64) //Parse komutu da kullanılabilir.
	if err != nil {
		fmt.Println("Casting Hatası")
	} else {
		fmt.Println("Yaşınız :", yas)
	}
	fmt.Println(age)
}
