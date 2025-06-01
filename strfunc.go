package main

import "fmt"

func main() {
	//Menghitung panjang string
	fmt.Println(len("Andika Firansyah"))
	//Mengakses karakter pada string, kenapa outputnya 70?
	//Karena karakter pada string hasilnya adalah byte
	//sehingga harud dikonversi
	fmt.Println("Andika Firansyah"[7])
}