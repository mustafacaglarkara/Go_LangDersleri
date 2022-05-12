package main

import (
	"fmt"
	"runtime"
	"time"
)

func Selam(kisi string) {
	for i := 0; i < 1000; i++ {
		go fmt.Println("Merhabalar : ", kisi)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	//
	runtime.GOMAXPROCS(8)
	go Selam("World")
	Selam("Mustafa")
}
