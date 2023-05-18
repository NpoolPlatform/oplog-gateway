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
		oplog1.WithEntID(ctx, in.EntID),
		oplog1.WithNewValue(ctx, in.NewValue),
		oplog1.WithResult(ctx, in.Result),
		oplog1.WithFailReason(ctx, in.FailReason),
		oplog1.WithStatusCode(ctx, in.StatusCode),
		oplog1.WithReqHeaders(ctx, in.ReqHeaders),
		oplog1.WithRespHeaders(ctx, in.RespHeaders),
		oplog1.WithElapsedMillisecs(ctx, in.GetElapsedMillisecs()),
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
