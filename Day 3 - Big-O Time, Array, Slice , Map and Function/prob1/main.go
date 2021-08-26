package main

import (
	"fmt"
)

func primeNumber(number int) string {
	if number <= 1 {
		return "Bukan Bilangan Prima"
	}
	if number <= 3 {
		return "Bilangan Prima"
	}

	if number%2 == 0 || number%3 == 0 {
		return "Bukan Bilangan Prima"
	}

	for i := 5; i*i <= number; i = i + 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return "Bukan Bilangan Prima"
		}
	}

	return "Bilangan Prima"
}

func main() {

	fmt.Println(primeNumber(7)) // true

	fmt.Println(primeNumber(1500450271)) // true

	fmt.Println(primeNumber(1000000000)) // false

	fmt.Println(primeNumber(10000000019)) // true

	fmt.Println(primeNumber(10000000033)) // true

}
