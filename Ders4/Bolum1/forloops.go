package main

import "fmt"

func main() {
	// for looplar aşağıdaki gibi tanımlanır.
	//for i := 0; i < 100; i++ {
	//	fmt.Println(i)
	//}
	max := 100
	a, b := 0, 1
	for b < max {
		println(b)
		a, b = b, a+b
	}

	slc_1 := []int{1, 2, 3}
	for i := range [10]int{} {
		slc_1 = append(slc_1, i)
	}
	fmt.Println(slc_1)
}
