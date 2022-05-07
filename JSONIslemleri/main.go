package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	p1 := person{First: "Mustafa Çağlar", Last: "KARA", Age: 38, BildigiDiller: []programlamaDilleri{{Adi: "Kotlin", Seviye: 5}, {Adi: "GoLang", Seviye: 4}}}
	p2 := person{First: "Ülkü", Last: "KARA", Age: 38, BildigiDiller: []programlamaDilleri{{"C#", 5}, {"İngilizce", 5}}}

	people := []person{p1, p2}
	JsonData, err := json.MarshalIndent(people, "", " ")
	if err != nil {
		panic("HATA")
	} else {
		fmt.Println(string(JsonData))
	}

}

type person struct {
	First         string `json:"Adı"`
	Last          string `json:"Soyadı"`
	Age           int
	BildigiDiller []programlamaDilleri
}
type programlamaDilleri struct {
	Adi    string
	Seviye int
}
