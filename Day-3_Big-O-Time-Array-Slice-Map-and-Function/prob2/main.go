package main

import "fmt"

func pow(x, n int) int {

	// your code here
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	res := pow(x*x, n/2)
	if n%2 == 1 {
		res *= x
	}
	return res
}

func main() {

	fmt.Println(pow(2, 3)) // 8

	fmt.Println(pow(7, 2)) // 49

	fmt.Println(pow(10, 5)) // 100000

	fmt.Println(pow(17, 6)) // 24137569

	fmt.Println(pow(5, 3)) // 125

}
