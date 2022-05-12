package main

import "fmt"

func main() {
	myArray := [...]int{55, 28, 21, 52}
	//array üzerinde slice oluşturma
	mySlice := myArray[:]
	fmt.Println(mySlice)
	//slice a veri ekleme
	//mySlice = append(mySlice, 3355)
	fmt.Println(mySlice)
	// slice değişince arrayden de değişiyor
	mySlice[0] = 5555
	fmt.Println(myArray, mySlice)
	// slice dan veri alarak slice oluşturma
	denemeSlice := mySlice[:len(mySlice)-1]
	fmt.Println(denemeSlice)

	number := make([]int, 5, 5)
	number[0] = 12
	number[1] = 22
	number[2] = 33
	number[3] = 44
	number[4] = 55

	// kapazite artıyor
	number = append(number, 66)
	fmt.Println(number)
	fmt.Println(len(number), cap(number))

}
