package oplog

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/oplog/gw/v1/oplog"
	oplog1 "github.com/NpoolPlatform/oplog-gateway/pkg/oplog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateOpLog(ctx context.Context, in *npool.CreateOpLogRequest) (*npool.CreateOpLogResponse, error) {
	handler, err := oplog1.NewHandler(
		ctx,
		oplog1.WithAppID(ctx, in.GetAppID()),
		oplog1.WithUserID(ctx, in.UserID),
		oplog1.WithPath(ctx, in.GetPath()),
		oplog1.WithMethod(ctx, in.GetMethod()),
		oplog1.WithArguments(ctx, in.GetArguments()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOpLog",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOpLogResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateOpLog(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOpLog",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOpLogResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.CreateOpLogResponse{
		Info: info,
	}, nil
}
