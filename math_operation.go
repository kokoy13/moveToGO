package main

import "fmt"

func main() {
	a := 10
	b := 3
	c := a + b
	d := a - b
	e := a * b
	f := float32(a) / float32(b) // Pembagian harus dikonversi ke float32 untuk hasil desimal
	g := a % b // Modulus untuk mendapatkan sisa bagi
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("Hasil penjumlahan a + b:", c)
	fmt.Println("Hasil pengurangan a - b :", d)
	fmt.Println("Hasil perkalian a * b:", e)
	fmt.Println("Hasil pembagian a / b:", f)
	fmt.Println("Hasil modulus a % b :", g) // Menampilkan sisa bagi
}