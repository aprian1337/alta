package main

import (
	"fmt"
	"sort"
)

func PairSum(arr []int, target int) []int {

	// your code here
	sort.Ints(arr)
	left := 0
	right := len(arr) - 1
	for left < right {
		if arr[left]+arr[right] == target {
			return []int{left, right}
		} else if arr[left]+arr[right] < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}

func main() {

	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]

	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11)) // [0, 2]

	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12)) // [2, 3]

	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10)) // [1, 2]

	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6)) // [0, 1]

}
