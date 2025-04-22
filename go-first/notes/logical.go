package notes

import "fmt"

// logical operator test
func LogicalOperator() {
	a := true
	b := false
	c := 3
	d := 5

	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!a)

	fmt.Println(c & d)
	fmt.Println(c | d)
	fmt.Println(c ^ d)
	fmt.Println(^c)
	fmt.Println(c &^ d)
}
