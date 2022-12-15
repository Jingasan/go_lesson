// lesson35: sync.Mutexの練習
// 2つのgoroutineから1つのmapを読んだりすると発生するエラーを回避するためにsync.Mutexを使用する
// Mutexを用いることで、片方の処理が書き込み中の場合にもう片方が書き込めないようにロックする。
// 読み込みの場合も書き込みの場合と同様に制御する。
package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := Counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("Key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("Key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c, c.Value("Key"))
}
