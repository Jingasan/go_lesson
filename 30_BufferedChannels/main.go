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

	// Channelからの値取り出し
	x := <-ch
	fmt.Println(x)
	x = <-ch
	fmt.Println(x)

	// rangeによる取り出し
	ch <- 300
	fmt.Println(len(ch))
	ch <- 400
	fmt.Println(len(ch))
	close(ch) // rangeで取り出す場合は一度Channelをcloseする
	for c := range ch {
		fmt.Println(c)
	}
}
