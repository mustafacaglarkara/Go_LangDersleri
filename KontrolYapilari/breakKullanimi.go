package main

import "fmt"

func main() {
	count := 0
	var bolunebilenler = []int{}

	for i := 0; true; i++ {
		if i%13 == 0 && i >= 13 {
			count++
			fmt.Println("Şuanki bölünebilen sayı : ", i)
			bolunebilenler = append(bolunebilenler, i)
		}
		if count == 10 {
			break
		}
	}
	fmt.Println(bolunebilenler)
}
