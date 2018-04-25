package filesystem

import (
	"os"
	"strings"
)

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func GetDirSimpleName(dir string) string {
	parts := strings.Split(dir, "/")
	return parts[len(parts)-1]
}
