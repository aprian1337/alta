package main

import "fmt"

func helper(n int, Cache map[int]int) int {
	i, ok := Cache[n]
	if ok {
		return i
	}
	Cache[n] = helper(n-1, Cache) + helper(n-2, Cache)
	return Cache[n]
}

func fibo(n int) int {
	cache := make(map[int]int)
	cache[0] = 0
	cache[1] = 1
	return helper(n, cache)
}

func main() {

	//fmt.Println(fibo(0)) // 0
	//
	//fmt.Println(fibo(1)) // 1
	//
	//fmt.Println(fibo(2)) // 1
	//
	//fmt.Println(fibo(3)) // 2
	//
	//fmt.Println(fibo(5)) // 5
	//
	//fmt.Println(fibo(6)) // 8
	//
	//fmt.Println(fibo(7)) // 13
	//
	//fmt.Println(fibo(9)) // 13

	fmt.Println(fibo(50)) // 55

}
