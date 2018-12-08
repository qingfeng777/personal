package utils

import (
	"fmt"
	"log"
)

func LogInfoF(label string, msg ...interface{}) string {
	for _, v := range msg {
		label = fmt.Sprintf(label+"   %v", v)
	}

	log.Println("[info] ", label)
	return label
}

func LogErrorF(label string, msg ...interface{}) string {
	for _, v := range msg {
		label = fmt.Sprintf(label+"   %v", v)
	}

	log.Println("[error] ", label)
	return label
}