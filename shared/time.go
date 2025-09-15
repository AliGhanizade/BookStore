package shared

import "time"

func GetCurrentTime() string {
	timeLayout := time.DateTime
	return time.Now().Format(timeLayout)
}