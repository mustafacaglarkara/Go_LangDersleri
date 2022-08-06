package main

/*
func main() {
	data := "Sonuna Eklenecek olan veri"
	metin := []byte(data)
	err := os.WriteFile("okunacak.txt", metin, 0777)
	if err != nil {
		fmt.Println("Dosya yazma işlemi sırasında hata meydana geldi")
	}

}

func main() {
	dosya, hata := os.Open("okunacak1.txt")
	if hata != nil {
		fmt.Printf("Hata : %s \n", hata.Error())
		return
	}
	defer dosya.Close()
	DosyaBilgileri, hataBilgiler := dosya.Stat()
	if hataBilgiler != nil {
		fmt.Printf("Hata : %s ", hataBilgiler)
	}
	fmt.Println(DosyaBilgileri.IsDir())
}
/*
func main() {
	dosya, err := ioutil.ReadFile("okunacak.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dosya)
}

*/
