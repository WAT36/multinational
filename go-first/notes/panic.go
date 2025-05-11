package notes

import "fmt"

// rune test
func PanicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("リカバリしました:", r)
		}
	}()

	fmt.Println("処理開始")
	panic("パニック発生！")
	fmt.Println("ここは実行されません")
}
