package main

import "fmt"

func main() {
	var mesaj = "Hi"
	fmt.Println(add(10, 35))
	mesajDegistir(&mesaj)
	fmt.Println(mesaj)
}

func mesajDegistir(mesaj *string) {
	println(*mesaj)
	*mesaj = "Deneme Mesaj"
}

// x ve y parametre return ile iki parametrenin toplamı döndürülüyor.
func add(x, y int) int {
	return x + y
}
