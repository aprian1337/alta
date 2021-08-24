package main

import(
	"fmt"
)

const phi = 3.14

func main() {
	// input
	var T, r float64

	// kode disini
	fmt.Println("Masukkan nilai T : ")
	fmt.Scanf("%f\n", &T)
	fmt.Println("Masukkan nilai r : ")
	fmt.Scanf("%f\n", &r)
	Lp := 2*phi*(r*r) + 2*phi*r*T
	fmt.Println("Output : ", Lp)
	
}