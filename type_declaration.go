package main

import "fmt"

func main() {
	type NoKTP string

	// Deklarasi variabel dengan tipe data NoKTP
	var ktpKaka NoKTP = "1234567890123456"

	//Deklarasi variabel dengan tipe data string
	ktpWulan := "6543210987654321"
	//Konversi tipe data string ke NoKTP
	ktpWulanNoKTP := NoKTP(ktpWulan)
	fmt.Println("No KTP Kaka:", ktpKaka)
	fmt.Println("No KTP Wulan:", ktpWulanNoKTP)

}