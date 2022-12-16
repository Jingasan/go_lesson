// lesson36: packageの自作練習
package main

import (
	"fmt"

	"github.com/Jingasan/go_package/36_package/greeting"
	"github.com/Jingasan/go_package/36_package/sample"
)

func main() {
	fmt.Print(sample.Example())
	greeting.Hello()
	greeting.GoodMorning()
}
