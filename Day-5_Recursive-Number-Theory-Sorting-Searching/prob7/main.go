package main

import (
	"fmt"
)

func playingDomino(cards [][]int, deck []int) interface{} {

	// your code here
	max := 0
	var idxMax []int
	for i := 0; i < len(cards); i++ {
		if cards[i][0] == deck[0] || cards[i][0] == deck[1] || cards[i][1] == deck[0] || cards[i][1] == deck[1] {
			if max < cards[i][0]+cards[i][1] {
				max = cards[i][0] + cards[i][1]
				idxMax = cards[i]
			}
		}
	}
	if max != 0 {
		return idxMax
	}
	return "tutup kartu"
}

func main() {

	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))

	// [3, 4]

	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))

	// [6 5]

	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))

	// “tutup kartu”

}
