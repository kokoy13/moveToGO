package main

import "fmt"

func main() {
	//Nilai constant tidak dapat ditimpa atau diubah karena nilainya konstan atau tetap
	const phi = 3.14

	//multiple constant
	const(
		phi2 = 3.14
		phi3 = 22/7
	)

	fmt.Println("Nilai phi adalah", phi)
	fmt.Println("Nilai phi adalah", phi2, "atau sama dengan", phi3)
}