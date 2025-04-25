package notes

import (
	"fmt"
)

// closure
func Closure() func() int {
	a := 0

	return func() int {
		a++
		return a
	}
}

func main() {
	c := Closure()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
}
