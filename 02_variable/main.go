// 変数，定数の学習
package main

import "fmt"

var i int = 1

const Pi = 3.14
const (
	Username = "test_user"
	Password = "test_pass"
)

func main() {
	// 変数
	var t, f bool = true, false
	var (
		f64   float64 = 1.2
		s, sr string  = "test", "Yes"
	)
	fmt.Println(i, f64, s, sr, t, f)
	yi := 1
	yf64 := 1.2
	xs := "test"
	xt := true
	fmt.Println(yi, yf64, xs, xt)

	// 定数
	fmt.Println(Pi, Username, Password)

	// 変数の型
	var (
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)
	fmt.Println(u8, i8, f32, c64)
	fmt.Printf("type=%T value=%v", u8, u8)
}
