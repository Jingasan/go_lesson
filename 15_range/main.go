// lesson15: rangeの練習
package main

import (
	"fmt"
)

func main() {
	// 配列でのrange
	l := []string{"python", "go", "java"}
	for i, v := range l {
		fmt.Println(i, v)
	}
	// 配列でのrange（インデックス番号の省略）
	for _, v := range l {
		fmt.Println(v)
	}
	// mapでのrange
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
	// mapでのrange（Keyのみ）
	for k := range m {
		fmt.Println(k)
	}
	// mapでのrange（Valueのみ）
	for _, v := range m {
		fmt.Println(v)
	}
}
