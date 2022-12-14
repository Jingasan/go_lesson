// lesson8: スライス(可変長配列)の練習
package main

import (
	"fmt"
)

func main() {
	// スライス
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])
	fmt.Println(n[:])
	n[2] = 200
	fmt.Println(n)

	// 要素の追加
	n = append(n, 100, 200, 300, 400)
	fmt.Println(n)

	// 二次元配列
	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board)
}
