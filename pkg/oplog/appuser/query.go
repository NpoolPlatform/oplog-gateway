//nolint
package appuser

import (
	"context"
	"fmt"
	// "encoding/json"
	"path/filepath"
	"strings"
	// appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	// npool "github.com/NpoolPlatform/message/npool/appuser/gw/v1/app"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func GetEntity(ctx context.Context, path, args string) (interface{}, error) {
	strs := strings.Split("path", "/")
	if len(strs) < 4 {
		return nil, fmt.Errorf("invalid path")
	}
	if !strings.HasPrefix(strs[3], "update") {
		return nil, nil
	}

	_path := filepath.Join(strs[3:]...)
	switch _path {
	case "/update/app":
		/*
			var req npool.UpdateAppRequest
			if err := json.Unmarshal([]byte(args), &req); err != nil {
				return nil, err
			}
			return appmwcli.GetApp(ctx, req.ID)
		*/
	default:
	}

	return nil, nil
}

func HumanReadable(path, args, cur, new string, result basetypes.Result, reason string) (string, error) {
	return "", nil
}
