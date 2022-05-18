package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(TahminEt(0))
	fmt.Println(TahminEt(101))
	fmt.Println(TahminEt(50))
	mesaj, hata := TahminEt(-25)
	fmt.Println(mesaj, hata)

}
func TahminEt(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("Hata 1-100 arasında bir sayı giriniz!")
	}
	return "Bildiniz", nil
}
