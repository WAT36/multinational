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

	// エラーハンドリング例
	err := func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("エラー発生: %v", r)
			}
		}()
		// ここでパニックを発生させる
		panic("パニック発生！")
		// ここは実行されません
		return nil
	}()

	if err != nil {
		fmt.Println("エラーを検知:", err)
	} else {
		fmt.Println("正常終了")
	}
}
