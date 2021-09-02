package main

import "fmt"

func DragonOfLoowater(dragonHead, knightHeight []int) {
	// your code here
	sortKnightHeight := MergeSort(knightHeight)
	sortDragonHead := MergeSort(dragonHead)
	if sortKnightHeight[len(sortKnightHeight)-1] > sortDragonHead[len(sortDragonHead)-1] && sortKnightHeight[len(sortKnightHeight)-2] > sortDragonHead[len(sortDragonHead)-2] {
		for i, v := range sortKnightHeight {
			if v >= sortDragonHead[0] {
				for j := i + 1; j < len(sortKnightHeight); j++ {
					if sortKnightHeight[j] >= sortDragonHead[1] {
						fmt.Println(v + sortKnightHeight[j])
						return
					}
				}
			}
		}
	} else {
		fmt.Println("knight fall")
	}

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

	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4, 3}) // 11

	DragonOfLoowater([]int{5, 10}, []int{5}) // knight fall

	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall

	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10

}
