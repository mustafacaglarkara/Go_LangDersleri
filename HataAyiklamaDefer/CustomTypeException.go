package main

import (
	"fmt"
)

type BorderException struct {
	parameter int
	message   string
}

func (b *BorderException) Error() string {
	return fmt.Sprintf("%d -> %s ", b.parameter, b.message)
}

func TahminEt2(tahmin int) (string, error) {
	if tahmin <= 1 || tahmin > 100 {
		return "", &BorderException{tahmin, "Sayı 1 ile 100 arasında olmalıdır."}
	}
	return "Bildiniz", nil

}
func main() {
	sonuc, err := TahminEt2(0)
	fmt.Println(sonuc, err)
}
