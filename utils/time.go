package utils

import "time"

func CurrentMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
