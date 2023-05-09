package oplog

import (
	"fmt"
	"path/filepath"
	"strings"
)

func getPathPrefix(path string) (string, error) {
	strs := strings.Split(path, "/")
	if len(strs) < 3 {
		return "", fmt.Errorf("invalid path")
	}
	if !strings.HasPrefix(strs[2], "v") {
		return "", fmt.Errorf("invalid path")
	}
	return filepath.Join(strs[0], strs[1]), nil
}
