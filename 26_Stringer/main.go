// lesson26: Stringer（Stringの出力方法を変更したいときに利用）の練習
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	fmt.Println("My name is " + p.Name + ".")
	return fmt.Sprintf("My name is %v.", p.Name)
}

func main() {
	mike := Person{"Mike", 22}
	fmt.Println(mike)
}
