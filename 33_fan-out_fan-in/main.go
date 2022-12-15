// lesson33: fan-out, fan-inの練習
// goroutineを使用して結果を渡すことをfan-out、受け取ることをfan-inという
package main

import "fmt"

func producer(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

func second(first <-chan int, second chan<- int) {
	defer close(second)
	for i := range first {
		second <- i * 2
	}
}

func third(second chan int, third chan int) {
	defer close(third)
	for i := range second {
		third <- i + 10000
	}
}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan int)

	go producer(chan1)
	go second(chan1, chan2)
	go third(chan2, chan3)
	for result := range chan3 {
		fmt.Println(result)
	}
}
