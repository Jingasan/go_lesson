// lesson36: packageの自作練習
package main

import (
	"36_package/greeting"
	"36_package/sample"
	"fmt"
)

func main() {
	fmt.Print(sample.Example())
	greeting.Hello()
	greeting.GoodMorning()
}
