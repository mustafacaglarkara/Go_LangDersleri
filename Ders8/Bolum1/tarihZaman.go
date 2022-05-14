package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2022, time.May, 13, 15, 55, 0, 0, time.UTC)
	fmt.Println(t)
	// mevcut zaman
	fmt.Println(time.Now())
	// Ay bilgisi
	fmt.Println(time.Now().Month())
	yarin := time.Now().AddDate(0, 0, 1)
	fmt.Println(yarin)

	longformat := "Friday, May 2, 2006"
	fmt.Println("Bugün ", yarin.Format(longformat))
	shortFormat := "1/2/2006"
	fmt.Println("Bugün : ", yarin.Format(shortFormat))

	eskiTarih := time.Date(1983, time.July, 10, 0, 0, 0, 0, time.UTC)
	fmt.Println(eskiTarih.Weekday())
}
