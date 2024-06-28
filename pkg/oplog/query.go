package oplog

import (
	"context"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/pubsub"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	oplogmwpb "github.com/NpoolPlatform/message/npool/oplog/mw/v1/oplog"
	appuser "github.com/NpoolPlatform/oplog-gateway/pkg/oplog/appuser"
	oplogmwcli "github.com/NpoolPlatform/oplog-middleware/pkg/client/oplog"
)

type queryHandler struct {
	*Handler
	infos []*oplogmwpb.OpLog
	reqs  []*oplogmwpb.OpLogReq
}

func (h *queryHandler) updateHumanReadable() {
	if len(h.reqs) == 0 {
		return
	}

	if err := pubsub.WithPublisher(func(publisher *pubsub.Publisher) error {
		return publisher.Update(
			basetypes.MsgID_UpdateOpLogHumanReadableReq.String(),
			nil,
			nil,
			nil,
			h.reqs,
		)
	}); err != nil {
		logger.Sugar().Errorw(
			"updateHumanReadable",
			"AppID", h.AppID,
			"Error", err,
		)
	}
}

func (h *queryHandler) formalizeHumanReadable() error {
	for _, info := range h.infos {
		if info.HumanReadable != "" {
			continue
		}

		req := &oplogmwpb.OpLogReq{
			EntID: &info.EntID,
		}

		prefix, err := getPathPrefix(info.Path)
		if err != nil {
			return wlog.WrapError(err)
		}
		var str string

		switch prefix {
		case "/api/appuser":
			str, err = appuser.HumanReadable(
				info.Path,
				info.Arguments,
				info.CurValue,
				info.NewValue,
				info.Result,
				info.FailReason,
			)
		case "/api/notif":
		default:
		}
		if err != nil {
			return wlog.WrapError(err)
		}

		req.HumanReadable = &str
		h.reqs = append(h.reqs, req)
	}
	return nil
}

func (h *Handler) GetOpLogs(ctx context.Context) ([]*oplogmwpb.OpLog, uint32, error) {
	conds := &oplogmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
	}
	if h.UserID != nil {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID}
	}

	infos, total, err := oplogmwcli.GetOpLogs(ctx, conds, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler := &queryHandler{
		Handler: h,
		infos:   infos,
	}
	if err := handler.formalizeHumanReadable(); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	handler.updateHumanReadable()

	return handler.infos, total, nil
}
