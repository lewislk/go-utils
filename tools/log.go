package tools

import "encoding/json"

func GetLogStr[T any](input T) string {
	log, _ := json.Marshal(input)
	return string(log)
}
