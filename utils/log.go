package utils

import (
	"fmt"
	"log"
)

func LogInfo(message string, args ...interface{}) {
	log.Println(fmt.Sprintf(message, args...))
}
