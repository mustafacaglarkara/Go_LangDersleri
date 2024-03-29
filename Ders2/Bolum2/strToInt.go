package main

import (
	"fmt"
	"strconv"
)

func main() {
	var myString = "1"
	var x = 10
	var f float32 = 2.0
	var s1 = "3.123"

	// String datayı float a çevirme
	f1, _ := strconv.ParseFloat(s1, 64)
	fmt.Println(f1)

	fmt.Println(myString, x, f)

	//string to int
	// dönüşüm sırasında hata olması durumunda _ yerine err konulabilir
	number, _ := strconv.Atoi(myString)
	// strconv.PArseInt ile convert işlemi
	ikincoConvert, _ := strconv.ParseInt(myString, 10, 64)
	fmt.Println(number, ikincoConvert)
	// int to str convert işlemi
	fmt.Println("Sonuç : ", strconv.Itoa(x))
	//float to int convert
	fmt.Println(int(f))
}
