// lesson8: スライス(可変長配列)の練習
package main

import (
	"fmt"
)

func main() {
	// スライスの宣言
	fmt.Println("スライスの宣言")
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
	fmt.Println("要素の追加")
	n = append(n, 100, 200, 300, 400)
	fmt.Println(n)

	// 二次元配列
	fmt.Println("二次元配列")
	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board)

	// makeによるスライスのキャパシティ（要素数上限）の指定
	fmt.Println("makeによるスライスのキャパシティ（要素数上限）の指定")
	m := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(m), cap(m), m)
	m = append(m, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(m), cap(m), m)
	m = append(m, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(m), cap(m), m)
	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)
	b := make([]int, 0)
	var c []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)
}
