package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("İsminizi Giriniz :")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Println("Yaşınızı Giriniz : ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Yaşınız : ", f)
	}
}
