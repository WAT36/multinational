package notes

import "fmt"

func Channel() {
	ch := make(chan int, 10)

	// チャネルに送信
	ch <- 5

	// チャネルから受信
	i := <-ch

	fmt.Println(i)
}
