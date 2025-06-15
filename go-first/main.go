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

// go run main.go
func main() {
	utils.PrintMessage()

	//notes.Rune()
	//notes.LogicalOperator()
	//closure()
	//notes.Defer()
	//notes.PanicRecover()
	//notes.Go()
	//notes.Slice()
	//notes.Channel()
	//notes.Pointer()
	// notes.Method()
	//notes.Tag()
	notes.InterfaceTest()
}
