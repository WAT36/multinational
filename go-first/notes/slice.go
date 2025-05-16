package notes

import "fmt"

// defer test
func Slice() {
	s := make([]int, 10)
	fmt.Println(s)
	s[0] = 9
	fmt.Println(s)
}
