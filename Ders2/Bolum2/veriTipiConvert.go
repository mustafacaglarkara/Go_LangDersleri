package main

import (
	"fmt"
	"strconv"
)

//func main() {
//	var age int
//	var adiniz string
//	fmt.Println("Lütfen yaşınızı giriniz")
//	fmt.Scanf("%d", &age) // & işareti ile memorideki age veri gönderdik
//	fmt.Println("Girilen yaş :=> ", age)
//	// bu komut ile ekrana girilen yaş yazar.
//	//ancak örnek 38 yazdıktan sonra string bir değer girer iseniz
//	// onu kout olarak algılar ve hata verir. komut bulunamadı şeklinde.
//
//	fmt.Println("Lütfen adınızı giriniz")
//	fmt.Scanf("%s", &adiniz)
//	fmt.Println("Merhabalar adınız : ", adiniz)
//	// ama burda da bir sıkıntı var. girilen değerde boşluk olursa yine onu komut olarak algılar
//}

func main() {
	var input string
	fmt.Println("Lütfen yaşınızı giriniz")
	fmt.Scanf("%s", &input)
	age, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Yaşınız : ", age)
	}
}
