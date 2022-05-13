package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	runtime.GOMAXPROCS(8)
	go printX()
	go printY()
	go printZ()
	wg.Wait()
	fmt.Println("Yazma İşlemi Bitti")
}
func printX() {
	for i := 0; i <= 100000; i++ {
		fmt.Println(i, "-->", "X")
	}
	wg.Done()
}
func printY() {
	for i := 0; i <= 100000; i++ {
		fmt.Println(i, "-->", "Y")
	}
	wg.Done()
}
func printZ() {
	for i := 0; i <= 100000; i++ {
		fmt.Println(i, "-->", "Z")
	}
	wg.Done()
}
