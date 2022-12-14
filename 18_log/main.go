// lesson18: logの練習
package main

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func main() {
	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	LoggingSettings("test.log")
	_, err := os.Open("fdafdsafa")
	if err != nil {
		log.Fatalln("Exit", err)
	}
}
