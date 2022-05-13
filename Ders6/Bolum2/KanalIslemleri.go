package main

import (
	"fmt"
)

func main() {
	// int tipinde slice
	sayilar := []int{1, 2, 3, 4, 5, 6}
	// int tipinde channel oluşturma
	xChan := make(chan int)
	// Goroutine parametre olarak fonksiyonda gerekli olan
	//int ve channel tipinde parametre
	go toplama(sayilar, xChan)
	// fonksiyondan gelen değeri x değişkenine atama
	x := <-xChan
	fmt.Println(x)
}

// goroutine ile çağırdığımız fonksiyon
func toplama(sayi []int, ch chan int) {
	sum := 0
	for _, say := range sayi {
		sum += say
	}
	ch <- sum
}
