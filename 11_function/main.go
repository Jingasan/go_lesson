// lesson11: 関数の練習
package main

import (
	"fmt"
)

func calculate(x, y int) (int, int) {
	fmt.Println("calculate function")
	return x + y, x - y
}

func mul(price, item int) (result int) {
	fmt.Println("multiply function")
	result = price * item
	return
}

func add(x int, y int) int {
	fmt.Println("add function")
	return x + y
}

func foo(params ...int) {
	fmt.Println(len(params), params)
}

func main() {
	// 関数のコール
	r := add(10, 20)
	fmt.Println(r)
	r1, r2 := calculate(10, 20)
	fmt.Println(r1, r2)
	fmt.Println(mul(100, 2))

	// 関数の内部で関数を使用1
	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)
	// 関数の内部で関数を使用2
	func(x int) {
		fmt.Println("inner func", x)
	}(2)

	// 可変長引数
	foo()
	foo(10, 20)
	foo(10, 20, 30)

	// スライス渡し
	s := []int{1, 2, 3}
	foo(s...)

	// 外部ソースコードの関数呼び出し
	Bar()
}
