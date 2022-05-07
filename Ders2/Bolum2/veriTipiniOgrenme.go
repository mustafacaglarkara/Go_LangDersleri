package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	firstName, lastName, age := "Mustafa Çağlar", "KARA", 38

	start := time.Now()
	// Printf ile kullanımı
	fmt.Printf("%T start değişken tipi \n ", start)         // time.Time
	fmt.Printf("%T firstName değişken tipi \n ", firstName) //string
	fmt.Printf("%T lastName değişken tipi \n ", lastName)   //string
	fmt.Printf("%T age değişken tipi \n ", age)             //int

	// Println ile kullanımı
	fmt.Println(" Start değişkeni", reflect.TypeOf(start))
	fmt.Println(" Start değişkeni", reflect.ValueOf(start).Kind())

}
