package main

import "fmt"

func cetakTablePerkalian(number int) {

	// Process: Your Solution Code Here
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}
}

func main() {

	cetakTablePerkalian(35)

}
