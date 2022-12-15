// lesson30: Buffered Channelの練習
package main

import "fmt"

func main() {
	// Channelには指定した数だけ値を追加できる
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))
}
