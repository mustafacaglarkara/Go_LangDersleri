package main

import "fmt"

type Insan struct {
	Isim    string
	Soyisim string
	Yas     int
}

func (insan Insan) IsımYaz() {
	fmt.Printf("Adı: %s soyadı: %s yaşı %d \n", insan.Isim, insan.Soyisim, insan.Yas)
}
func (insan *Insan) IsımEkle(yeniIsim string) {
	insan.Isim = yeniIsim
}
func main() {

	i := Insan{Isim: "Mustafa Çağlar", Soyisim: "KARA", Yas: 38}
	i.IsımYaz()
	i.IsımEkle("Ülkü")
	i.IsımYaz()
	//isimler := []string{"Mustafa", "Çağlar"}
	//
	//isimler = append(isimler, "Ülkü", "Öykü")
	//buyukler := []string{"Yılmaz", "Şükran"}
	//fmt.Println(isimler)
	//isimler = append(isimler, buyukler...)
	//fmt.Println(isimler)
	//sonuc := sort.SearchStrings(isimler, "Mustafa")
	//isimler = append(isimler[:sonuc], isimler[sonuc+1:]...)
	//fmt.Println(isimler)
}
