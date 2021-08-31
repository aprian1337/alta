package main

import "fmt"

func BinarySearch(array []int, x int) {

	// your code here
	sortArr := MergeSort(array)
	startIndex := 0
	endIndex := len(sortArr) - 1
	midIndex := len(sortArr) / 2
	for startIndex <= endIndex {

		value := sortArr[midIndex]

		if value == x {
			fmt.Println(midIndex)
			return
		}

		if value > x {
			endIndex = midIndex - 1
			midIndex = (startIndex + endIndex) / 2
			continue
		}

		startIndex = midIndex + 1
		midIndex = (startIndex + endIndex) / 2
	}
	fmt.Println(-1)
}

func Merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}

func MergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	n := len(s) / 2
	l := MergeSort(s[:n])
	r := MergeSort(s[n:])
	return Merge(l, r)
}

func main() {

	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3) // 2

	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5) // 3

	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53) // 6

	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) // -1

}
