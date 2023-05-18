package oplog

import (
	"context"

	oplogmwpb "github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"
	oplogmwcli "github.com/NpoolPlatform/oplog-middleware/pkg/client/oplog"
)

func (h *Handler) UpdateOpLog(ctx context.Context) (*oplogmwpb.OpLog, error) {
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
