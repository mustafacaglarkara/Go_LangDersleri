package main

import "fmt"

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

// Contructer işlemi Varsayılan ve Boş Yapımı Metod
func NewHuman() *Human {
	h := new(Human)
	return h
}

//Parametreli Constructor
/*
Metod overload Go da yok o yüzden isim değiştirilmesi
gerekmektedir.
*/
func NewHuman2(firstname, lastname string, age int) *Human {
	h := new(Human)
	h.FirstName = firstname
	h.LastName = lastname
	h.Age = age
	return h
}

func main() {
	// birinci yöntem
	insan := Human{}
	insan.FirstName = "Mustata Çağlar"
	//insan.LastName = "KARA"
	insan.Age = 38

	fmt.Println(insan)

	// ikinci tanımala
	erkek := new(Human)
	erkek.LastName = "KARA"
	fmt.Println(erkek)
	fmt.Println(erkek.LastName)

	//Yapıcı Metod ile Kullanımı
	x := NewHuman()
	x.Age = 38
	fmt.Println(x.Age)

	//Yapıcı Metod2 Kullanımı
	den := NewHuman2("Ülkü", "KARA", 38)
	fmt.Println(*den)
}
