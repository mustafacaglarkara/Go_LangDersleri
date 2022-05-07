package main

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
}
