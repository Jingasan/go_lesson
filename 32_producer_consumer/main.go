// lesson32: producerとconsumerの練習
// producer
//   様々な処理を並列で行う処理
//   (ex) 色々なサーバーに行き、logを取ってきて解析したりする
//   処理が終わったらchannelに値を送信する
// consumer
//   channelから受け取った値を処理する
package main

import (
	"fmt"
	"sync"
)

func producer1(ch chan int, i int) {
	ch <- i * 2
}

func producer2(ch chan int, i int) {
	ch <- i*2 + 1000
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("process", i)
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup
	producerNum := 2
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(producerNum)
		go producer1(ch, i)
		go producer2(ch, i)
	}

	go consumer(ch, &wg)
	wg.Wait()
	close(ch)
}
