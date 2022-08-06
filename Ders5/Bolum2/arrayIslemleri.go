package main

import "fmt"

func main() {
	//Basit Dizi
	myArray := [3]int{}
	myArray[0] = 32
	myArray[1] = 23
	myArray[2] = 54
	fmt.Println(myArray)
	fmt.Printf("%#v \n ", myArray) // çıktı olarak [3]int{32,23,54} gelir

	// Renk Dizisi
	color := [3]string{}
	color[0] = "Red"
	color[1] = "Green"
	color[2] = "Blue"

	fmt.Println(color[1]) //get
	color[2] = "White"    //set
	fmt.Println(color[2])
	fmt.Println(color)

	// Direk tanımala
	sayilarim := [3]int{27, 12, 83}
	fmt.Println(sayilarim)

	fmt.Println(len(sayilarim))

	//Belirtmeden tanımlama

	var myArrayGenis = [...]int{5, 8, 97, 8, 9, 22}
	myArrayGenis[5] = 55
	fmt.Println(myArrayGenis)

	// out of range hatası verir
	//var myArrayGenis2 = [...]int{}
	//myArrayGenis2[5] = 50

	//string array listesi
	var myCars = [3]string{"Volvo", "Mercedes", "Renault"}
	fmt.Println(myCars, len(myCars), cap(myCars))
	for _, car := range myCars {
		fmt.Println(car)
	}
	// index ile veri getirme
	for i, _ := range myCars {
		fmt.Println(myCars[i])
	}

}
