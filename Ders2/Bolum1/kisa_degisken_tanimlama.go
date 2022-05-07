package main

import "fmt"

//func main() {
//	firstName := "Mustafa" //kısa değişken tanımala
//	firstName, lastName, age := "Mustafa Çağlar", "KARA", 38 //noktalı virgül ile hızlıca değişken tanımalarız.
//	var firstName, lastName, age = "Mustafa Çağlar", "KARA", 38 // var ile tanımlanınca iki nokta koymuyoruz.
//	fmt.Println(firstName, lastName, age)
//}

//Hatalı tanımlama
//func main(){
//	var firstName, lastName string, age int = "Mustafa Çağlar", "KARA", 38 // hata verir
//	var firstName,lastName string = "Mustafa","KARA" // düzgün tanımalama
//}

// Parantez içinde değişken tanımala
func main() {
	//var firstName, lastName string, age int = "Mustafa Çağlar", "KARA", 38  hata vermiş olduğundan farklı veri tipleri ile ilgili tanımalama
	// yaparken alttaki metodu kullanabilirsiniz.
	var (
		firstName string = "Mustafa Çağlar"
		lastName  string = "KARA"
		age       int    = 25
	)
	fmt.Println(firstName, lastName, age)
}
