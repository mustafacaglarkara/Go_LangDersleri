package main

import "fmt"

func main() {
	kare := kare{kenar: 10}
	sekil(kare)
}

type shape interface {
	area() float64
}

type kare struct {
	kenar float64
}

func (k kare) area() float64 {
	return k.kenar * 4
}

func sekil(s shape) {
	fmt.Println(s.area())
}
