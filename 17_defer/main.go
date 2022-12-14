// lesson17: deferによる遅延実行の練習
package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")
	fmt.Println("hello foo")
}

func main() {
	// deferの基礎
	defer fmt.Println("world")
	foo()
	fmt.Println("hello")

	// deferの使いどころ：ファイルクローズの予約
	file, _ := os.Open("./main.go")
	defer file.Close()
	data := make([]byte, 50)
	file.Read(data)
	fmt.Println(string(data))
}
