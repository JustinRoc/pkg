package util

import (
	"encoding/json"
)


func JSONLog(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	return string(b)
}