package main

import "fmt"

func main() {
	// string değişken tanımlama
	var mesaj string = "Merhabalar"

	// değişken tipi tanımala
	var ikinciMesaj string
	// değişkene değer verme
	ikinciMesaj = "Nasılsın"

	// değişken tipi verilmeden tanımlama
	// fonksiyon içerisinde geçerli
	ucuncuMesaj := "Fonksiyon içerisinde geçerli"

	//toplu değişken tanımlama
	// aynı veri tipinde
	var birinci, ikinci, ucuncu string = "Birinci", "İkinci", "Üçüncü"

	// atama yapılmadı. ama derleme aşamasında hata vermez
	// default değer olan 0 yazılır.
	var a, b, c int

	// atama yapılmadı yazdırılınca false default olan değier yazar
	var durum bool

	// tip belirtmeden değişlen tanımlama, sıralı bir şekilde
	//var adi, yasi, mDurumu = "Mustafa Çağlar KARA", 38, true
	// fonksiyon içi kısa değişken tanımala
	adi, yasi, mDurumu := "Mustafa Çağlar KARA", 38, true

	fmt.Println(mesaj)
	fmt.Println(ikinciMesaj)
	fmt.Println(ucuncuMesaj)
	fmt.Println(birinci, ikinci, ucuncu)
	fmt.Println(a, b, c)
	fmt.Println(durum)
	fmt.Println(adi, yasi, mDurumu)
}
