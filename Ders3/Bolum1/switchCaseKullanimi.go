package main

import "fmt"

func main() {
	sayi := 1
	haftaninGunu := ""

	switch sayi {
	case 1:
		haftaninGunu = "Pazartesi"
		fmt.Println("Çok zor tatilden sonra... :)")
	case 2:
		haftaninGunu = "Salı"
	case 3:
		haftaninGunu = "Çarşamba"
	case 4:
		haftaninGunu = "Perşembe"
	case 5:
		haftaninGunu = "Cuma"
	case 6:
		haftaninGunu = "Cumartesi"
	case 7:
		haftaninGunu = "Pazar"
	default:
		haftaninGunu = "--Hata--"
	}
	fmt.Println(haftaninGunu)
}
