package main

import (
	"fmt"
	"time"
)

func main() {
	//Unix zaman 1 Ocak 1970 den beri geçen zaman
	suankiUnixZaman := time.Now().Unix()
	fmt.Println(suankiUnixZaman)
	time.Sleep(time.Second)
	suankiUnixZaman = time.Now().Unix()
	fmt.Println(suankiUnixZaman)

}
