package main

import "fmt"

var sehirler = make(map[string]int)

func main() {
	sehirler["Samsun"] = 55
	sehirler["Giresun"] = 28
	sehirler["Ordu"] = 52
	sehirler["Diyarbakır"] = 21

	fmt.Println(sehirler)
	fmt.Println(sehirler["Samsun"])
	// Sehirler Mapinden Samsunu Siler
	delete(sehirler, "Samsun")
	fmt.Println(sehirler)

	// map içinde veri olup olmadığı kontrölü
	if data, durum := sehirler["Ordu"]; durum {
		fmt.Println(data)
		fmt.Println(durum)
	} else {
		fmt.Println("Kayıt Bulunamadı")
	}
	fmt.Println(mapIcindekiVeriBul("Samsun"))
}

// map içinde veri arayan fonksion
func mapIcindekiVeriBul(veri string) bool {
	if _, durum := sehirler[veri]; durum {
		return durum
	} else {
		return durum
	}
}
