package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

func main() {

	// デッドロック再現用プログラム

	// 複数のgoroutineの完了を待つための値
	var wg sync.WaitGroup

	printSum := func(v1, v2 *value) {

		// wgをデクリメント
		defer wg.Done()

		// v1
		v1.mu.Lock()
		defer v1.mu.Unlock()

		// 処理の負荷をシュミレートする為、一定時間スリープし
		// デッドロックを誘発する
		time.Sleep(2 * time.Second)

		//v2
		v2.mu.Lock()

		defer v1.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)

	}

	var a, b value
	// wgをインクリメント
	wg.Add(2)

	// goroutine 実行
	go printSum(&a, &b)
	go printSum(&b, &a)

	// wg.Addで追加したすべてgoroutine が、Doneで終了通知されるまで待機
	wg.Wait()

}
