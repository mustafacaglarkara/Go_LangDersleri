package main

import "fmt"

var ilkIsimler = []string{"Mustafa", "Çağlar", "Ülkü"}
var ikinciIsimler = make([]string, 2)

func main() {
	fmt.Println(ikinciIsimler)
	copy(ikinciIsimler, ilkIsimler[:2])
	fmt.Println(ikinciIsimler, "", ilkIsimler)
}
