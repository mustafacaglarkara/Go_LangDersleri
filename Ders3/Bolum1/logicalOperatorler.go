package main

import "fmt"

//func main() {
//	num := 6
//	condition := num > 2 && num < 9
//	fmt.Println(condition) // true
//}

func main() {
	var username string
	fmt.Println("Kullanıcı Adını Giriniz")
	fmt.Scanf("%s", &username)
	if username != "mck" {
		fmt.Println("Kullanıcı Adı Hatalı")
	} else {
		fmt.Println("Hoş geldin admin.")
	}
}
