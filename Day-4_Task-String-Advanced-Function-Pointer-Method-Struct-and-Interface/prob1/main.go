package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {

	// your code here
	if strings.Compare(a, b) == 1 {
		return b
	} else if strings.Compare(b, a) == 1 {
		return a
	} else {
		return "No matches string found"
	}
}

func main() {
	fmt.Println(Compare("AKA", "ZJJAKASHI"))  // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
