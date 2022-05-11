package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{10, 15}
	v.Y = 55
	v.X = 10
	p := &v
	fmt.Println(*p)
	fmt.Println(v)

}
