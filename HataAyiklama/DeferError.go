package main

import "fmt"

func main() {
	fmt.Println("Program Başladı")
	KareAl(5)
}

func KareAl(sayi int) {

	defer func() {
		hata := recover()
		if hata != nil {
			fmt.Println(hata)
			fmt.Println("Hata Meydana Geldi")
		}
	}()
	panic("--- Program Durdu ---")
	fmt.Println(sayi * sayi)
}
