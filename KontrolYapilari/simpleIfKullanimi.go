package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if sayi, err := strconv.Atoi("45"); err == nil {
		fmt.Println(sayi)
	} else {
		fmt.Println("Geçersiz Sayı Biçimi")
	}

	if args := os.Args; len(args) != 2 {
		fmt.Println("Argüman sayısı geçersiz.")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println(km)
	} else {
		fmt.Printf("%d km in miles is %v\n", km, float64(km)*0.621)
	}

}
