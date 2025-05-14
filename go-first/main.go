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

func init() {
	fmt.Println("Starting myapp...")
}

func init() {
	fmt.Println("Starting myapp2...")
}

func init() {
	fmt.Println("Starting myapp3...")
}

// go run main.go
func main() {
	utils.PrintMessage()

	//notes.Rune()
	//notes.LogicalOperator()
	//closure()
	//notes.Defer()
	//notes.PanicRecover()
	notes.Go()
}
