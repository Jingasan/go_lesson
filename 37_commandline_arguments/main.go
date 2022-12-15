// lesson37: コマンドライン引数の練習
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println("count: [", len(os.Args), "]")
	for i, v := range os.Args {
		fmt.Printf("args[%d] -> %s\n", i, v)
	}
}
