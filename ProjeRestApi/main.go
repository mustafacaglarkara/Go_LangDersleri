package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mck.com/goLandDersleri/ProjeRestApi/Entity"
	"net/http"
)

func main() {
	//var cat = Entity.Category{ID: 3, CategoryName: "Deneme Title"}
	//fmt.Println(cat)
	//var product = Entity.Product{ID: 4, ProductName: "Puding", CategoryID: 2, UnitPrice: 8.50}
	//fmt.Println(product)
	var products, err = GetAllProducts()
	if err != nil {
		fmt.Println(err)
	} else {
		if len(products) > 0 {
			for _, urun := range products {
				fmt.Println(urun.ProductName)
			}
		} else {
			fmt.Println("Ürün Bulunamadı")
		}
	}

	AddProduct()
}

func GetAllProducts() ([]Entity.Product, error) {
	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	byteProducts, errRead := ioutil.ReadAll(response.Body)
	if errRead != nil {
		return nil, err
	}
	//fmt.Println(string(byteProducts))
	var productList []Entity.Product
	json.Unmarshal(byteProducts, &productList)
	//fmt.Println(productList)
	return productList, nil
}
func AddProduct() (Entity.Product, error) {
	//Url adresi girilir
	var urlAdres = "http://localhost:3000/products"
	// gönderilecek olan veri set edilir.
	product := Entity.Product{CategoryID: 1, ProductName: "Mac Pro", UnitPrice: 34000}
	// struct to json
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		return Entity.Product{}, err
	}
	// Http post işlemi
	response, errResponse := http.Post(urlAdres, "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	if errResponse != nil {
		return Entity.Product{}, nil
	}
	// bağlantıyı sonlarıracak olan func
	defer response.Body.Close()
	// gonderme işlemi err yok ise tmdır. gelen cevabı alıcaz.
	fmt.Println("Kaydedildi.")
	bodyByte, _ := ioutil.ReadAll(response.Body)
	//bodyString := string(bodyByte)
	//fmt.Println(bodyString)
	//todo struct çevirme
	var productResponse Entity.Product
	json.Unmarshal(bodyByte, &productResponse)
	//fmt.Println("Kayıt Edildi", productResponse.ProductName)
	return productResponse, nil
}
