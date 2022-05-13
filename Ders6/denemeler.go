package main

import "fmt"

func main() {
	sayiListesi := []int{50, 12, 38, 98, 78, 100}
	xChan := make(chan int)
	go toplamaIslemi(sayiListesi, xChan)
	go toplamaIslemi(sayiListesi[:len(sayiListesi)-2], xChan)
	x, y := <-xChan, <-xChan
	fmt.Println(x)
	fmt.Println(y)
}

func toplamaIslemi(sayilar []int, c chan int) {
	toplam := 0
	for _, deger := range sayilar {
		toplam += deger
	}
	c <- toplam
}
