package notes

import (
	"fmt"
	"time"
)

func say(message string) {
	for i := 0; i < 3; i++ {
		fmt.Println(message)
		time.Sleep(100 * time.Millisecond)
	}
}

func Go() {
	go say("こんにちは") // 並行処理
	say("Hello")    // メインスレッド
}
