// lesson13: if elseの練習
package main

import (
	"fmt"
)

func main() {
	s := 9
	if s%2 == 0 {
		fmt.Println("by 2")
	} else if s%3 == 0 {
		fmt.Println("by 3")
	} else {
		fmt.Println("else")
	}

	fmt.Println("==========")

	x, y := 10, 20
	v, w := 20, 30
	if x == 10 && y == 20 {
		fmt.Println("and1")
	}

	fmt.Println("==========")

	if v == 20 && w == 20 {
		fmt.Println("and2")
	}

	fmt.Println("==========")

	if x == 10 || y == 10 {
		fmt.Println("or1")
	}

	fmt.Println("==========")

	if v == 30 || w == 20 {
		fmt.Println("or2")
	}
}
