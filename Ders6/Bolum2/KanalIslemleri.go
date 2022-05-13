package main

import (
	"fmt"
)

func main() {
	sayilar := []int{1, 2, 3, 4, 5, 6}
	xChan := make(chan int)
	go toplama(sayilar, xChan)
	x := <-xChan
	fmt.Println(x)
}

func toplama(sayi []int, ch chan int) {
	sum := 0
	for _, say := range sayi {
		sum += say
	}
	ch <- sum
}
