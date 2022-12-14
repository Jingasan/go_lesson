// lesson14: forの練習
package main

import (
	"fmt"
)

func main() {
	// for文の基本構文
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}
		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}
	l := []string{"python", "go", "java"}
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}
	// 省略形
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	// 無限ループ
	for {
		fmt.Println("endless")
	}
}
