package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {

	// your code here
	isTrue := map[string]int{}
	var res []int
	for _, v := range angka {
		if isTrue[string(v)] == 0 {
			isTrue[string(v)] = 1
		} else {
			isTrue[string(v)]++
		}
	}

	for k, v := range isTrue {
		if v == 1 {
			conv, _ := strconv.Atoi(k)
			res = append(res, conv)
		}
	}

	return res
}

func main() {

	fmt.Println(munculSekali("1234123")) // [4]

	fmt.Println(munculSekali("76523752")) // [6, 3]

	fmt.Println(munculSekali("12345")) // [1 2 3 4 5]

	fmt.Println(munculSekali("1122334455")) // []

	fmt.Println(munculSekali("0872504")) // [8 7 2 5 4]

}
