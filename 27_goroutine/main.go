// lesson27: goroutineによる並列処理の練習
package main

import (
	"fmt"
	"time"
)

// 並列処理関数
func goroutine(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go goroutine("world")
	normal("hello")
}
