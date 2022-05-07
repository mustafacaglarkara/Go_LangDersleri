package main

import "fmt"

func doSomething() (int, bool) {
	return 5, true
}

func main() {
	value, error := doSomething()
	if error {
		fmt.Println("Hata Var")
	} else {
		fmt.Println("Hata Yok gelen deger : ", value)
	}

}
