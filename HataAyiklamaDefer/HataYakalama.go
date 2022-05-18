package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("deneme.txt")
	if err != nil {
		//Type Assertion
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dizin Hatalı", pErr.Path)
			return
		} else {
			fmt.Println("Dosya Bulunamadı", pErr.Err)
			return
		}

	} else {
		fmt.Println(file.Name())
	}

}
