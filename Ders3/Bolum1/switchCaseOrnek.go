package main

import "fmt"

func main() {
	puan := 88
	not := ""
	switch {
	case puan >= 90:
		not = "A"
	case puan >= 80:
		not = "B"
	case puan >= 70:
		not = "C"
	case puan >= 60:
		not = "D"
	default:
		not = "F"
	}
	fmt.Println(not)
}
