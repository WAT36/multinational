package notes

import "fmt"

func Pointer() {
	i := 5
	ip := &i
	fmt.Printf("%p", ip)
}
