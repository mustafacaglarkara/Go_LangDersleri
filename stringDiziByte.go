package main

import (
	"fmt"
)

func main() {
	var isim string = "Mustafa Çağlar KARA"
	var bIsım []rune = []rune(isim)
	fmt.Println(bIsım)

	for bIndex, bVeri := range bIsım {
		fmt.Println(bVeri, " -- ", string(bIsım[bIndex]))
	}
}
