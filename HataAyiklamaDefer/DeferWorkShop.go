package main

import (
	"fmt"
)

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	defer LogKayıt(p)
	fmt.Println("Ürün Başarılı Bir şekilde Kayıt Edildi.", p.productName)
}
func LogKayıt(s product) {
	fmt.Println(s.productName, " isimli ürün kayıt işlemi gerçekleşti.")
}
func main() {
	p1 := product{productName: "Oppo Reno 5", unitPrice: 4500}
	p1.save()
}
