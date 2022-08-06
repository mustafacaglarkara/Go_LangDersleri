package main

import (
	"fmt"
	"os"
)

func main() {
	dosya, hata := os.Open("okunacak.txt")
	if hata != nil {
		fmt.Println("Dosya Bulunamadı", hata)
		return
	}
	defer dosya.Close()
	// Dosya bilgisi alınıyor
	dosyaBilgisi, _ := dosya.Stat()
	//Dosya boyutu kadar byte oluşturuluyor.
	dosyaKesit := make([]byte, dosyaBilgisi.Size())

	//Okunan byte lar string e çeviriliyor
	okunanByte, _ := dosya.Read(dosyaKesit)

	okunanMetin := string(dosyaKesit)

	fmt.Println("Okunan Byte :", okunanByte)
	fmt.Println("Okunan Metin : \n", okunanMetin)

}
