// lesson7: 配列の練習
package main

import (
	"fmt"
)

func main() {
	// 配列の宣言
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)
	// 配列の初期化
	var b [2]int = [2]int{100, 200}
	fmt.Println(b)
}
