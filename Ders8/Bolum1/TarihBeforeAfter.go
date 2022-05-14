package main

import (
	"fmt"
	"time"
)

func main() {

	oAn := time.Date(1983, time.July, 10, 0, 0, 0, 0, time.UTC)
	suan := time.Now()

	fmt.Println(oAn.Before(suan)) // oncesi mi
	fmt.Println(oAn.After(suan))  // sonrasi mi
	fmt.Println(oAn.Equal(suan))  // eşit mi?

	//	Dateif Kullanımı
	fmt.Println(suan.Sub(oAn))
	fmt.Println(suan.Sub(oAn).Hours())
}
