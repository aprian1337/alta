package main

import (
	"fmt"
	_ "fmt"
)

func MaximumBuyProduct(money int, productPrice []int) {

	// your code here
	insertionSort(&productPrice)
	count := 0
	for i := range productPrice {
		if money-productPrice[i] >= 0 {
			money -= productPrice[i]
			count++
		} else {
			fmt.Println(count)
			return
		}
	}
}

func insertionSort(num *[]int) {
	for i := 0; i < len(*num); i++ {
		tmp := (*num)[i]
		j := i
		for j > 0 && (*num)[j-1] > tmp {
			(*num)[j] = (*num)[j-1]
			j--
		}
		(*num)[j] = tmp
	}
}

func main() {

	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000}) // 3

	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}) // 4

	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000}) // 4

	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000}) // 1

	MaximumBuyProduct(0, []int{10000, 30000}) // 0

}
