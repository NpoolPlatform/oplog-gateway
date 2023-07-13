package oplog

import (
	"context"
	"encoding/json"
	"fmt"

	oplogmwpb "github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"
	oplogmwcli "github.com/NpoolPlatform/oplog-middleware/pkg/client/oplog"
)

func (h *Handler) UpdateOpLog(ctx context.Context) (*oplogmwpb.OpLog, error) {
	if h.EntID == nil {
		return nil, fmt.Errorf("invalid entid")
	}

	const maxNewValueLen = 256
	if h.NewValue != nil && len(*h.NewValue) > maxNewValueLen {
		var val struct {
			Infos []interface{}
		}
		if err := json.Unmarshal([]byte(*h.NewValue), &val); err == nil {
			value := fmt.Sprintf(`{"Infos":{"Count":%v}}`, len(val.Infos))
			h.NewValue = &value
		}
	}

	return oplogmwcli.UpdateOpLog(ctx, &oplogmwpb.OpLogReq{
		EntID:            h.EntID,
		NewValue:         h.NewValue,
		Result:           h.Result,
		FailReason:       h.FailReason,
		StatusCode:       h.StatusCode,
		ReqHeaders:       h.ReqHeaders,
		RespHeaders:      h.RespHeaders,
		ElapsedMillisecs: h.ElapsedMillisecs,
	})
}
