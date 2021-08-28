package main

import (
	"fmt"
	"sort"
)

type pair struct {
	key   string
	count int
}

func MostAppearItem(items []string) []pair {

	// your code here
	itemWithMap := map[string]int{}
	for _, v := range items {
		if _, ok := itemWithMap[v]; ok {
			itemWithMap[v]++
		} else {
			itemWithMap[v] = 1
		}
	}

	var pairs []pair
	for k, v := range itemWithMap {
		pairs = append(pairs, pair{key: k, count: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count < pairs[j].count
	})
	return pairs
}

func main() {

	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))

	//golang->1 ruby->2 js->4

	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))

	// C->1 D->2 B->3 A->4

	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))

	//football->1 basketball->1 tenis->1

}
