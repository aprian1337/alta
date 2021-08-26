package main
import "fmt"

func palindrome(input string) bool {
	// your code here
	idxLast := len(input)-1
	idxStart := 0
	for idxStart!=len(input) {
		if input[idxLast] != input[idxStart]{
			return false
		}
		idxLast--
		idxStart++
	}
	return true
}

func main() {

	fmt.Println(palindrome("civic"))       // true

	fmt.Println(palindrome("katak"))       // true

	fmt.Println(palindrome("kasur rusak")) // true

	fmt.Println(palindrome("mistar"))      // false

	fmt.Println(palindrome("lion"))        // false

}