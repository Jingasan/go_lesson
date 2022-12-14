// lesson22: 構造体の練習
package main

import "fmt"

type Vertex struct {
	X, Y int
	S    string
}

func main() {
	// 構造体の基本
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 100
	fmt.Println(v.X, v.Y)
	v2 := Vertex{X: 1}
	fmt.Println(v2)
	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)
	v4 := Vertex{}
	fmt.Println(v4)
	// 構造体のポインタ
	var v5 Vertex
	fmt.Printf("%T %v\n", v5, v5)
	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6)
	v7 := &Vertex{} // 構造体の場合、ポインタはnewよりも&がよく使われる
	fmt.Printf("%T %v\n", v7, v7)
}
