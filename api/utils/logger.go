package utils

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "INFO: ", log.LstdFlags)

func LogInfo(message string) {
	logger.Println(message)
}

func LogError(message string, err error) {
	logger.Printf("ERROR: %s - %v", message, err)
}
