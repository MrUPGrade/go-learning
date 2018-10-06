package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Program is starting")

	programName := filepath.Base(os.Args[0])

	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}

	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!")
	//log.Fatal(sysLog)
	log.Panic(sysLog)
	fmt.Println("Program is ending")
}
