// 論理値型の練習
package main

import "fmt"

func main() {
	// 論理値型
	t, f := true, false
	fmt.Printf("%T %v %t\n", t, t, 1)
	fmt.Printf("%T %v %t\n", f, f, 5)
	fmt.Println("===========")

	// and, or, not
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)
	fmt.Println("===========")
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)
	fmt.Println("===========")
	fmt.Println(!true)
	fmt.Println(!false)
}
