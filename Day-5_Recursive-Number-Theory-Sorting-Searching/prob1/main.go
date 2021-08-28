package main

import (
	"fmt"
	"math"
)

func primeX(number int) int {
	// your code here
	if number == 1 {
		return 2
	} else if number == 2 {
		return 3
	}
	var prime int
	check := 5
	number -= 2
	for number != 0 {
		if getPrime(check) == true {
			prime = check
			number--
		}
		check++
	}
	return prime
}

func getPrime(check int) bool {
	for i := 2; i <= int(math.Sqrt(float64(check))); i++ {
		if check%i == 0 {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29

}
