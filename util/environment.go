package util

import (
	"os"
	"strconv"
)

func getEVString(key string, defaultValue string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return defaultValue
	}

	return value
}

func getEVInt64(key string, defaultValue int64) int64 {
	value := os.Getenv(key)

	if len(value) == 0 {
		return defaultValue
	}

	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultValue
	}

	return result
}

func EVString(key string, defaultValue string) string {
	return getEVString(key, defaultValue)
}

func EVInt(key string, defaultValue int) int {
	return int(getEVInt64(key, int64(defaultValue)))
}

func EVInt64(key string, defaultValue int64) int64 {
	return getEVInt64(key, defaultValue)
}
