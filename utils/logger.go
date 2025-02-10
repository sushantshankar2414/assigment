package utils

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "[LOG] ", log.LstdFlags)

func LogInfo(message string) {
	Logger.Println("INFO:", message)
}

func LogError(err error) {
	Logger.Println("ERROR:", err)
}