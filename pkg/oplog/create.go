package oplog

import (
	"context"
	"encoding/json"

	oplogmwpb "github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"
	oplogmwcli "github.com/NpoolPlatform/oplog-middleware/pkg/client/oplog"

	appuser "github.com/NpoolPlatform/oplog-gateway/pkg/oplog/appuser"
)

func (h *Handler) CreateOpLog(ctx context.Context) (*oplogmwpb.OpLog, error) {
	prefix, err := h.GetPathPrefix(ctx)
	if err != nil {
		return nil, err
	}

	valAndErr := map[string]interface{}{}
	var curVal interface{}
	curValStr := "{}"

	switch prefix {
	case "/api/appuser":
		curVal, err = appuser.GetEntity(ctx, h.Path, h.Arguments)
	case "/api/notif":
	default:
	}
	if err != nil {
		valAndErr["Error"] = err
	}
	if curVal != nil {
		valAndErr["Info"] = curVal
	}

	_curVal, err := json.Marshal(valAndErr)
	if err != nil {
		return nil, err
	}
	curValStr = string(_curVal)

	return oplogmwcli.CreateOpLog(ctx, &oplogmwpb.OpLogReq{
		AppID:     &h.AppID,
		UserID:    h.UserID,
		Path:      &h.Path,
		Method:    &h.Method,
		Arguments: &h.Arguments,
		CurValue:  &curValStr,
	})
}
