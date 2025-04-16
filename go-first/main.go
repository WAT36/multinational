package main

import (
	"fmt"
	"go-first/notes" // notes パッケージのインポート
	"go-first/utils" // utils パッケージのインポート
)

// go run main.go
func main() {
    fmt.Println("Starting myapp...")
    utils.PrintMessage()

	// rune test
	notes.Rune()
}
