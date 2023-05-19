package helper

import (
	"errors"
	"os"
	"strconv"
)

// EnvOrDefault returns the value of the environment variable k or defaultVal if it is not set.
func EnvOrDefault(k string, defaultVal string) string {
	v := os.Getenv(k)
	if v == "" {
		return defaultVal
	}
	return v
}

// EnvOrDefaultInt64 returns the value of the environment variable k or defaultVal if it is not set.
// It panics if the value is not an int64.
// Note: Panics are not recoverable. Use this function only in main() or init() functions.
func EnvOrDefaultInt64(k string, defaultVal int64) int64 {
	v := os.Getenv(k)
	if v == "" {
		return defaultVal
	}
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(errors.New("environment variable " + k + " is not an int64"))
	}
	return i
}

// GetEnv returns the value of the environment variable k or if it is not set an error.
func GetEnv(k string) (string, error) {
	if v, ok := os.LookupEnv(k); ok {
		return v, nil
	}

	return "", errors.New("environment variable " + k + " not set")
}

// MustGetEnv returns the value of the environment variable k or panics if it is not set.
// Note: Panics are not recoverable. Use this function only in main() or init() functions.
// Otherwise, use GetEnv to handle the error.
func MustGetEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		panic("environment variable " + k + " not set")
	}
	return v
}

// Short returns a shortened string of length 7.
// Useful for shortening commit hashes.
// And nothing else.
// Really.
func Short(s string) string {
	if len(s) > 7 {
		return s[0:7]
	}
	return s
}
