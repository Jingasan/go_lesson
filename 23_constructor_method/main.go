// lesson23: コンストラクタとメソッドの練習
package main

import "fmt"

type Vertex struct {
	X, Y int
}

// 関数
func Area(v Vertex) int {
	return v.X * v.Y
}

// メソッド：ポインタレシーバー（ポインタを渡している関数
// -> 大元の構造体の値が変わる（大元の構造体の値を変えたいときに使う）
func (v *Vertex) ScaleP(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

// メソッド：値レシーバー（値を渡している関数）
// -> 大元の構造体の値が変わらない
func (v Vertex) ScaleV(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

// メソッド：値レシーバー（値を渡している関数）
// -> 大元の構造体の値が変わらない
func (v Vertex) Area() int {
	return v.X * v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Area(v))
	v.ScaleV(10)
	fmt.Println(v.Area())
	v.ScaleP(10)
	fmt.Println(v.Area())
}
