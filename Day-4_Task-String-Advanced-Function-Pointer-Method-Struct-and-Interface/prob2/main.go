package main

import (
	"fmt"
)

func caesar(offset int, input string) string {

	// your code here
	result := []rune(input)
	for k, v := range input {
		temp := v + rune(offset%26)
		if temp <= 122 {
			result[k] = temp
		} else if temp <= 148 {
			result[k] = temp - 26
		}
	}
	return string(result)
}

func main() {

	fmt.Println(caesar(3, "abc")) // def

	fmt.Println(caesar(2, "alta")) // cnvc

	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi

	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))

	//bcdefghijklmnopqrstuvwxyza

	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))

	// mnopqrstuvwxyzabcdefghijkl

}
