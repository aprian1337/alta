
package main

import "fmt"

func main() {

	// input
	var studentScore = 64

	// Process: Your Solution Code Here
	if studentScore < 35 && studentScore >= 0{
		fmt.Println("E")
	} else if studentScore < 50 {
		fmt.Println("D")
	} else if studentScore < 65 {
		fmt.Println("C")
	} else if studentScore < 80 {
		fmt.Println("B")
	} else if studentScore <= 100 {
		fmt.Println("A")
	}else {
		fmt.Println("Nilai Invalid")
	}

}