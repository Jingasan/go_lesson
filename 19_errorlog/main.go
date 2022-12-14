// lesson19: Fatalfの練習
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		// Fatalfを使うことでerrorlogを出力でき、その後の処理は実行されない
		log.Fatalln("Error!")
	}
	defer file.Close()
	data := make([]byte, 20)
	count, err := file.Read(data)
	if err != nil {
		// Fatalfを使うことでerrorlogを出力でき、その後の処理は実行されない
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))

	// Fatalfを使うことでerrorlogを出力でき、その後の処理は実行されない
	log.Fatalf("%T %v", "test", "test")
	fmt.Println("ok?")
}
