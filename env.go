package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func String(key string, value interface{}) string {
	res := os.Getenv(key)
	switch res {
	case "":
		switch value.(type) {
		case string:
			return value.(string)
		case []byte:
			return string(value.([]byte))
		default:
			return fmt.Sprint(value)
		}
	default:
		return res
	}
}

func Strings(key string, delim string, defaultValue ...string) []string {
	res := strings.Split(os.Getenv(key), delim)
	switch len(res) == 1 && res[0] == "" {
	case true:
		return defaultValue
	default:
		return res
	}
}

func Int(key string, defaultValue interface{}) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		switch defaultValue.(type) {
		case int:
			return defaultValue.(int)
		case int8:
			return int(defaultValue.(int8))
		case int16:
			return int(defaultValue.(int16))
		case int32:
			return int(defaultValue.(int32))
		case int64:
			return int(defaultValue.(int64))
		case uint:
			return int(defaultValue.(uint))
		case uint8:
			return int(defaultValue.(uint8))
		case uint16:
			return int(defaultValue.(uint16))
		case uint32:
			return int(defaultValue.(uint32))
		case uint64:
			return int(defaultValue.(uint64))
		default:
			return 0
		}
	}
	return value
}

func Float(key string, defaultValue interface{}) float64 {
	i, err := strconv.ParseFloat(os.Getenv(key), 64)
	if err != nil {
		switch defaultValue.(type) {
		case float64:
			return defaultValue.(float64)
		case int:
			return float64(defaultValue.(int))
		case int8:
			return float64(defaultValue.(int8))
		case int16:
			return float64(defaultValue.(int16))
		case int32:
			return float64(defaultValue.(int32))
		case int64:
			return float64(defaultValue.(int64))
		case uint:
			return float64(defaultValue.(uint))
		case uint8:
			return float64(defaultValue.(uint8))
		case uint16:
			return float64(defaultValue.(uint16))
		case uint32:
			return float64(defaultValue.(uint32))
		case uint64:
			return float64(defaultValue.(uint64))
		default:
			return 0
		}
	}
	return i
}

func Bool(key string, defaultValue interface{}) bool {
	switch os.Getenv(key) {
	case "true":
		return true
	case "false":
		return false
	case "1":
		return true
	default:
		switch defaultValue.(type) {
		case bool:
			return defaultValue.(bool)
		case int:
			return defaultValue.(int) > 0
		default:
			return false
		}
	}
}
