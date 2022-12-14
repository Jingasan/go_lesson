// lesson19: Fatalfの練習
package main

import (
	"fmt"
	"log"
)

func main() {
	// Fatalfを使うことでerrorlogを出力でき、その後の処理は実行されない
	log.Fatalf("%T %v", "test", "test")
	fmt.Println("ok?")
}
