package main

import "fmt"

func main() {
	insanlar := [5]string{"Mustafa Çağlar", "Ülkü", "Öykü Beren", "Şükran", "Yılmaz"}
	arkadaslar := [2]string{"Ülkü", "Beril"}
	count := 0
outher:
	for index, name := range insanlar {
		for _, friend := range arkadaslar {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outher
			}
		}
	}
	count++
	fmt.Println("Next instruction after the break")
	if count < 3 {
		goto outher
	}
}
