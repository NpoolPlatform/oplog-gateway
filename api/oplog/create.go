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
		oplog1.WithAppID(in.GetAppID(), true),
		oplog1.WithUserID(in.UserID, false),
		oplog1.WithPath(in.GetPath(), true),
		oplog1.WithMethod(in.GetMethod(), true),
		oplog1.WithArguments(in.GetArguments(), false),
		oplog1.WithReqHeaders(&in.ReqHeaders, false),
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
