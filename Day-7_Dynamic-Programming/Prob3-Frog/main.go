package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {

	// your code here
	var res []int
	//res = append(res, 0)
	res[0] = 0
	for i := 1; i < len(jumps); i++ {
		if i >= 2 {
			tmp1 := res[i-1] + int(math.Abs(float64(jumps[i]-jumps[i-1])))
			tmp2 := res[i-2] + int(math.Abs(float64(jumps[i]-jumps[i-2])))
			if tmp1 < tmp2 {
				res = append(res, tmp1)
			} else {
				res = append(res, tmp2)
			}
		} else {
			res = append(res, res[i-1]+int(math.Abs(float64(jumps[i]-jumps[i-1])))) //20
		}
	}
	return res[len(res)-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
