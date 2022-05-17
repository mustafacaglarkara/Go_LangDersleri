package main

import "fmt"

func main() {

	sozluk := make(map[string]int)
	sozluk2 := map[string]string{"bool": "Kitap"}

	fmt.Println(sozluk2)

	sozluk["Samsun"] = 55
	sozluk["Giresun"] = 28
	sozluk["Diyarbakır"] = 21

	fmt.Println(sozluk)
	delete(sozluk, "Diyarbakır")
	fmt.Println(sozluk)

	for key, value := range sozluk {
		fmt.Println(key, "-> ", value)
	}
}
