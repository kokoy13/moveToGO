package main

import "fmt"

func main() {
	// Deklarasi variabel
	var name string = "Andika Firansyah"

	// Deklarasi variabel tanpa tipe data(ini bisa dan direkomendasikan)
	var kota_lahir = "Padang"

	//bisa tidak menggunakan var jika menginisialisasi menggunakan :=
	//namun, jika jenis kelamin ditimpa ulang dibawahnya maka gunakan jenis_kelamin = "nilai" saja
	jenis_kelamin := "Laki-laki"

	var age int = 20
	var isBool bool = true

	fmt.Println("Nama saya", name)
	fmt.Println("Umur saya", age)

	name = "Andika Ganteng"

	fmt.Println("Tapi saya berubah pikiran, nama saya sekarang", name)
	fmt.Println("Umur saya masih", age)
	fmt.Println("Jenis kelamin saya", jenis_kelamin)
	fmt.Println("Kota lahir saya", kota_lahir)
	fmt.Println("Apakah betul?", isBool)
}