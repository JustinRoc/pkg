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

func truncateStringField(data any, path []string) {
	if len(path) == 0 || data == nil {
		return
	}

	switch v := data.(type) {
	case map[string]any:
		if val, ok := v[path[0]]; ok {
			if len(path) == 1 {
				// 到达目标字段，处理字符串
				if str, ok := val.(string); ok {
					if !strings.HasPrefix(str, "http") && len(str) > 30 {
						v[path[0]] = str[:30]
					}
				}
			} else {
				truncateStringField(val, path[1:])
			}
		}
	case []any:
		for _, item := range v {
			truncateStringField(item, path)
		}
	}
}

type JSONOption func(jsonBytes []byte) ([]byte, error)

func WithoutFields(fields ...string) JSONOption {
	return func(jsonBytes []byte) ([]byte, error) {
		var m any
		if err := json.Unmarshal(jsonBytes, &m); err != nil {
			return nil, err
		}
		for _, f := range fields {
			parts := strings.Split(f, ".")
			deleteField(m, parts)
		}
		return json.Marshal(m)
	}
}

func TruncateStringFields(fields ...string) JSONOption {
	return func(jsonBytes []byte) ([]byte, error) {
		var m any
		if err := json.Unmarshal(jsonBytes, &m); err != nil {
			return nil, err
		}
		for _, field := range fields {
			parts := strings.Split(field, ".")
			truncateStringField(m, parts)
		}
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
