package main

import "fmt"

var sayilar = make([]int, 5, 10)

func main() {
	var diziSayi = []int{10, 55, 87, 88, 99, 12, 25, 66, 98, 24, 22, 35, 68, 78, 44, 78, 100, 125, 887, 778}
	sayilar = append(sayilar, diziSayi...)

	fmt.Println(sayilar)
	fmt.Println(len(sayilar))
	fmt.Println(cap(sayilar))

}
