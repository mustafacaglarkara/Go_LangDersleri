package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
		var isim string
		fmt.Println("Yapılacakları Yazınız... Çıkmak için QUIT yazın")
		fmt.Scanln(&isim)
		fmt.Println(isim)
		if isim = strings.ToUpper(isim); isim == "QUIT" {
			break
		}
	}

}
