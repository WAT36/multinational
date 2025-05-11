package main

import (
	"fmt"
	"go-first/notes" // notes パッケージのインポート
	"go-first/utils" // utils パッケージのインポート
)

func closure() {
	// logical operator test
	c := notes.Closure()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
}

// go run main.go
func main() {
	fmt.Println("Starting myapp...")
	utils.PrintMessage()

	//notes.Rune()
	//notes.LogicalOperator()
	//closure()
	//notes.Defer()
	notes.PanicRecover()
}
