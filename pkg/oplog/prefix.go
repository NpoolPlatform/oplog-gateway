package oplog

import (
	"fmt"
	"path/filepath"
	"strings"
)

func getPathPrefix(path string) (string, error) {
	const leastPathUnits = 3
	strs := strings.Split(path, "/")
	if len(strs) < leastPathUnits {
		return "", fmt.Errorf("invalid path")
	}
	if !strings.HasPrefix(strs[2], "v") {
		return "", fmt.Errorf("invalid path")
	}
	return filepath.Join(strs[0], strs[1]), nil
}
