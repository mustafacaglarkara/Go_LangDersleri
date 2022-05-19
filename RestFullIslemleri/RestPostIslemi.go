package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var urlAdres = "https://jsonplaceholder.typicode.com/todos"
	todo := Todo{1, 2, "Alışveriş yapılacak", false}
	todoJson, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
	}
	response, errPost := http.Post(urlAdres, "application/json;charset=utf-8", bytes.NewBuffer(todoJson))
	if errPost != nil {
		fmt.Println(errPost)
	}
	defer response.Body.Close()
	// response okumak için
	bodyByte, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyByte)
	fmt.Println(bodyString)
	//todo struct çevirme
	var todoResponse Todo
	json.Unmarshal(bodyByte, &todoResponse)
	fmt.Println(todoResponse)
}
