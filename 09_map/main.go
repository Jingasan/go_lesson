// lesson9: map(key-value型)の練習
package main

import (
	"fmt"
)

func main() {
	// mapの宣言
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)

	// 値の存在確認
	fmt.Println(m["nothing"])
	v, ok := m["apple"]
	fmt.Println(v, ok)
	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	// makeによるmapの初期化（makeを使わないとエラーになる）
	m2 := make(map[string]int)
	m2["note"] = 5000
	fmt.Println(m2)
}
