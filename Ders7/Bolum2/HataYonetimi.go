package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	_, err := os.Open("deneme.txt")
	errorCheck(err)
	//if err != nil {
	//	log.Println("Dosya yok veya okunamıyor.")
	//log.Fatal("Hata", err)
	//} else {
	//	fmt.Println(file)
	//}
	metin := fmt.Errorf("%v : Gönderilen Hata : %s", time.Now(), err)
	fmt.Println(metin.Error())
}

func errorCheck(err error) {
	if err != nil {
		fmt.Println("Hata : ", err.Error())
		os.Exit(1)
	}
}
