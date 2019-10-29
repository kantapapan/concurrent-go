package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// ライブロックの再現プログラム
// ライブロックとは並行操作を行っているけれど、
// その操作はプログラムの状態を全く進めていないプログラムのことです。
func main() {
	// 先にゴルーチンを用意して、一斉に実行したい場合はCondを使います。
	cadence := sync.NewCond(&sync.Mutex{})

	go func() {
		for range time.Tick(1 * time.Millisecond) {
			// 登録されているすべてを実行
			cadence.Broadcast()
		}

	}()

	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	// tryDir は ある人がある方向に動いて見て、うまく動けたかを返します。
	// 各方向は、その方向dirに動こうとしている人数で表されます。
	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %v", dirName)
		// 最初に、ある方向に動こうとしていることを、
		// その方向に動く人数を1増やす事で宣言します。
		atomic.AddInt32(dir, 1)
		// ライブロックの例を示すために、各人間は同じスピード、同じ歩調で動かなければいけません。
		// takeStepは全ての人間が同じ歩調で歩くのをシュミレートしています。
		takeStep()

		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprint(out, ". Success!")
			return true
		}

		takeStep()
		// ここで全ての人がこの方向に進めないと気づいて諦めます。
		// ここでこの人がこの方向に進めないと気づいて諦めます。
		// ここではその方向に動く人数を1減らすことで対応します。
		atomic.AddInt32(dir, -1)
		return false
	}

	var left, right int32

	tryLeft := func(out *bytes.Buffer) bool {
		return tryDir("left", &left, out)
	}

	tryRight := func(out *bytes.Buffer) bool {
		return tryDir("right", &right, out)
	}

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() {
			fmt.Println(out.String())
		}()
		defer walking.Done()
		fmt.Fprintf(&out, "%v is trying to scoot:", name)

		// このプログラムが終わるように、試行回数に作為的な上限を設けました。
		// ライブロックがあるプログラムでは、そのような上限がなく、ないが故に問題になるのです！
		for i := 0; i < 5; i++ {
			// まずある人が左に行こうとします。もし失敗したら右に行こうとします。
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
	}

	// この変数はプログラムが両方の人間がお互いにすれ違えるようになる、あるはすれ違うのを諦めるまで
	// 待つ方法を提供しています。
	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)

	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()

}
