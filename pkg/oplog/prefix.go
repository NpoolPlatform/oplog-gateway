package oplog

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
)

func (h *Handler) GetPathPrefix(ctx context.Context) (string, error) {
	strs := strings.Split(h.Path, "/")
	if len(strs) < 3 {
		return "", fmt.Errorf("invalid path")
	}
	if !strings.HasPrefix(strs[2], "v") {
		return "", fmt.Errorf("invalid path")
	}
	return filepath.Join(strs[0], strs[1]), nil
}
