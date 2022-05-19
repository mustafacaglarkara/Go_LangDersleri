package main

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

//func main() {
//	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer response.Body.Close()
//	// response okumak için
//	bodyByte, _ := ioutil.ReadAll(response.Body)
//	bodyString := string(bodyByte)
//	fmt.Println(bodyString)
//	//todo struct çevirme
//	var todo Todo
//	json.Unmarshal(bodyByte, &todo)
//	fmt.Println(todo)
//}
