package oplog

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/oplog/gw/v1/oplog"
	oplog1 "github.com/NpoolPlatform/oplog-gateway/pkg/oplog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateOpLog(ctx context.Context, in *npool.UpdateOpLogRequest) (*npool.UpdateOpLogResponse, error) {
	handler, err := oplog1.NewHandler(
		ctx,
		oplog1.WithID(&in.ID, true),
		oplog1.WithEntID(&in.EntID, true),
		oplog1.WithNewValue(in.NewValue, false),
		oplog1.WithResult(in.Result, false),
		oplog1.WithFailReason(in.FailReason, false),
		oplog1.WithStatusCode(in.StatusCode, false),
		oplog1.WithReqHeaders(in.ReqHeaders, false),
		oplog1.WithRespHeaders(in.RespHeaders, false),
		oplog1.WithElapsedMillisecs(in.GetElapsedMillisecs(), false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOpLog",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOpLogResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateOpLog(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOpLog",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOpLogResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.UpdateOpLogResponse{
		Info: info,
	}, nil
}
