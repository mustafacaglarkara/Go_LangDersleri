package main

import (
	"encoding/json"
	"fmt"
)

//
//func main() {
//	kisiler := []string{}
//	copyKisiler := make([]string, 8, 40)
//	kisiler = append(kisiler, "Mustafa", "Ülkü", "Beren")
//	sonuc := copy(copyKisiler, kisiler)
//	fmt.Println(reflect.TypeOf(kisiler))
//	fmt.Println(reflect.TypeOf(copyKisiler))
//	fmt.Println(sonuc)
//	fmt.Println(kisiler)
//	fmt.Println(copyKisiler)
//}

func main() {
	sozluk := make(map[string]string)
	sozluk["adi"] = "Mustafa Çağlar"
	sozluk["soyadi"] = "KARA"

	jsonVeri, _ := json.MarshalIndent(sozluk, "", " ")
	fmt.Println(string(jsonVeri))

}
