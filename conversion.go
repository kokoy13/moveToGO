package main

import "fmt"

func main() {
	// Konversi tipe data
	// Konversi dari int ke int64

	//ini adalah tipe data int32
	age := 20
	// Konversi ke int64
	// int64 adalah tipe data yang lebih besar dari int32
	fiveyearlater := int64(age+5)
	fmt.Println("Umur saya sekarang", age)
	fmt.Println("Umur saya 5 tahun lagi", fiveyearlater)

	//Mengambil karakter F dari string "Andika Firansyah"
	getF := "Andika Firansyah"[7]

	//hasil dari getF adalah byte
	// sehingga harus dikonversi ke string
	F := string(getF)
	fmt.Println("Karakter ke-7 dari string 'Andika Firansyah' adalah", F)

}