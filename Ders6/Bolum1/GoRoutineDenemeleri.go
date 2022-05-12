package main

import (
	"fmt"
	"time"
)

func selamVer(kisi string) {
	for i := 0; i < 100; i++ {
		fmt.Println("Merhabalar : ", kisi)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go selamVer("World")
	selamVer("Mustafa")
}
