package main

import "fmt"

func main() {
	var sayi1 = 55
	Demo1(&sayi1)
	fmt.Println(sayi1)
}

func Demo1(sayi *int) {
	*sayi = *sayi + 1
}
