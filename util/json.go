package util

import (
	"encoding/json"
	"strings"
)

// 删除 map 或数组中指定路径的字段
func deleteField(data any, path []string) {
	if len(path) == 0 {
		return
	}
	switch v := data.(type) {
	case map[string]any:
		if len(path) == 1 {
			delete(v, path[0])
			return
		}
		if next, ok := v[path[0]]; ok {
			deleteField(next, path[1:])
		}
	case []any:
		for _, item := range v {
			deleteField(item, path)
		}
	}
}

type JSONOption func(jsonBytes []byte) ([]byte, error)

func WithoutField(field string) JSONOption {
	return func(jsonBytes []byte) ([]byte, error) {
		var m any
		if err := json.Unmarshal(jsonBytes, &m); err != nil {
			return nil, err
		}
		parts := strings.Split(field, ".")
		deleteField(m, parts)
		return json.Marshal(m)
	}
}

func ToJSONStr(v any, opts ...JSONOption) string {
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	for _, opt := range opts {
		if b, err = opt(b); err != nil {
			return err.Error()
		}
	}
	return string(b)
}

func FromJSON[T any](data []byte) (T, error) {
	var v T
	err := json.Unmarshal(data, &v)
	if err != nil {
		return v, err
	}
	return v, nil
}
