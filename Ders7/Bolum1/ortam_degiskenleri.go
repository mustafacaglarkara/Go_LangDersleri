package main

import (
	"fmt"
	"os"
)

func main() {
	// ortam değişkenlerinin listesi
	//for _, env := range os.Environ() {
	//	fmt.Println(env)
	//}
	uName := os.Getenv("USERNAME")
	goRoot := os.Getenv("GOROOT")
	goPath := os.Getenv("GOPATH")
	homePath := os.Getenv("HOMEPATH")
	fmt.Println("Username :", uName)
	fmt.Println("GoRoot :", goRoot)
	fmt.Println("GoPath :", goPath)
	fmt.Println("HomePath :", homePath)
}
