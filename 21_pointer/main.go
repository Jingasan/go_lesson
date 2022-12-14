// lesson21: ポインタの練習
package main

import "fmt"

func one(x *int) {
	*x = 1
}

func main() {
	// ポインタの基礎
	var n int = 100
	fmt.Println(n)
	fmt.Println(&n)
	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)

	// 値の置き換え
	var m int = 100
	one(&m)
	fmt.Println(m)

	// new: ポインタを空の状態で宣言する場合に使用
	var p1 *int = new(int)
	fmt.Println(p1)
	fmt.Println(*p1)
	*p1++
	fmt.Println(*p1)
	var p2 *int
	fmt.Println(p2) // nil
}
