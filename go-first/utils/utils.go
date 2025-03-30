package utils

import "fmt"

// 公開される関数（大文字で始まる → 外部からアクセス可能）
func PrintMessage() {
    fmt.Println("Hello from utils package!")
}
