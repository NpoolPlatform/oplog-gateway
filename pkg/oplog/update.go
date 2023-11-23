package oplog

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	oplogmwpb "github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"
	oplogmwcli "github.com/NpoolPlatform/oplog-middleware/pkg/client/oplog"
)

func (h *Handler) UpdateOpLog(ctx context.Context) (*oplogmwpb.OpLog, error) {
	info, err := oplogmwcli.GetOpLogOnly(ctx, &oplogmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("oplog not exist")
	}

	const maxNewValueLen = 256
	if h.NewValue != nil && len(*h.NewValue) > maxNewValueLen {
		var val struct {
			Infos []interface{}
		}
		if err := json.Unmarshal([]byte(*h.NewValue), &val); err == nil && len(val.Infos) > 0 {
			value := fmt.Sprintf(`{"Infos":{"Count":%v}}`, len(val.Infos))
			h.NewValue = &value
		}
	}

	return oplogmwcli.UpdateOpLog(ctx, &oplogmwpb.OpLogReq{
		ID:               h.ID,
		NewValue:         h.NewValue,
		Result:           h.Result,
		FailReason:       h.FailReason,
		StatusCode:       h.StatusCode,
		ReqHeaders:       h.ReqHeaders,
		RespHeaders:      h.RespHeaders,
		ElapsedMillisecs: h.ElapsedMillisecs,
	})
}
