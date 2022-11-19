package utils

import (
	"strings"
	"time"
)

func GenerateId() string {
	currentTime := time.Now().UTC().Format("20060102150405.000000000")
	newCurrentTime := strings.ReplaceAll(currentTime, ".", "")

	return newCurrentTime
}
