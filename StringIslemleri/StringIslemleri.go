package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	isim := "Mustafa"
	fmt.Println(strings.Count(isim, "a"))
	fmt.Println(strings.Contains(isim, "a"))
	fmt.Println(strings.Index(isim, "a")) //indexini dönüyor. yok ise -1 döner
	fmt.Println(strings.ToLower(isim))
	fmt.Println(strings.ToUpper(isim))
	fmt.Println(strings.ToTitle(isim))

	//Son ek var mı ?
	fmt.Println(strings.HasSuffix(isim, "a"))
	// İlk ek parametredeki gibi mi ?
	fmt.Println(strings.HasPrefix(isim, "M"))

	// Join ile array içi veri birleştirme
	harfler := []string{"Ü", "l", "k", "ü"}
	fmt.Println(strings.Join(harfler, ""))

	// Replace ile karakter değiştirmek
	// -1 olur ise tüm karakterleri değiştirir
	fmt.Println(strings.Replace(isim, "a", "A", 1))
	fmt.Println(strings.Replace(isim, "a", "A", -1)) //-1 olursa tümünü değiştirir

	// Split ile metin parçalama
	hesap_no := "2022/05/18/17/24"
	hpNo_parcali := strings.Split(hesap_no, "/")
	fmt.Println(hpNo_parcali[0])

	// Repeat ile tekrarlama
	fmt.Println(strings.Repeat(isim, 10))
	// türkçe karakterler için
	tr := unicode.TurkishCase
	metin := strings.ToLowerSpecial(tr, "ADIM MUSTAFA ÇAĞLAR KARA GİRESUNDA YAŞIYORUM")
	println(metin)
}
