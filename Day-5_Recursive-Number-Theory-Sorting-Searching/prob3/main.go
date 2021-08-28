package main

import (
	"fmt"
	"math"
)

func primaSegiEmpat(high, wide, start int) {

	// your code here
	start++
	count := 0
	for i := 0; i < wide; i++ {
		for j := 0; j < high; j++ {
			isFound := false
			for isFound == false {
				if getPrime(start) == true {
					fmt.Print(start)
					fmt.Print("\t")
					count += start
					isFound = true
				}
				start++
			}
		}
		fmt.Println()
	}
	fmt.Println(count)
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

	primaSegiEmpat(2, 3, 13)

	/*

	   17 19

	   23 29

	   31 37

	   156

	*/

	primaSegiEmpat(5, 2, 1)

	/*

	   2  3  5  7 11

	   13 17 19 23 29

	   129

	*/

}
