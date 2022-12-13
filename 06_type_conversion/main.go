// lesson6: 型変換
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int -> float64 変換
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)
	// float64 -> int 変換
	var y float64 = 1.5
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)
	// 文字列 -> 数値型 変換
	var s string = "15"
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i)
}
