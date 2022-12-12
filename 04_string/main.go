// 文字列の学習
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 文字列の連結
	fmt.Println("Hello " + "World")

	// 文字の抽出
	fmt.Println(string("Hello World"[0]))

	// 文字列の置換
	var s string = "Hello World"
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)

	// 文字列の存在確認
	fmt.Println(strings.Contains(s, "World"))
}
