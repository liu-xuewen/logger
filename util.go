package logger

import "os"

// CheckPathExist CheckPathExist
func CheckPathExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return os.IsExist(err)
	}
	return true
}
