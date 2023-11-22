//nolint:dupl
package oplog

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/oplog/gw/v1/oplog"
	oplog1 "github.com/NpoolPlatform/oplog-gateway/pkg/oplog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetOpLogs(ctx context.Context, in *npool.GetOpLogsRequest) (*npool.GetOpLogsResponse, error) {
	handler, err := oplog1.NewHandler(
		ctx,
		oplog1.WithAppID(in.GetAppID(), true),
		oplog1.WithUserID(&in.UserID, true),
		oplog1.WithOffset(in.GetOffset()),
		oplog1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOpLogs",
			"In", in,
			"Error", err,
		)
		return &npool.GetOpLogsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetOpLogs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOpLogs",
			"In", in,
			"Error", err,
		)
		return &npool.GetOpLogsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.GetOpLogsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetAppOpLogs(ctx context.Context, in *npool.GetAppOpLogsRequest) (*npool.GetAppOpLogsResponse, error) {
	handler, err := oplog1.NewHandler(
		ctx,
		oplog1.WithAppID(in.GetAppID(), true),
		oplog1.WithOffset(in.GetOffset()),
		oplog1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppOpLogs",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppOpLogsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetOpLogs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppOpLogs",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppOpLogsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.GetAppOpLogsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetNAppOpLogs(ctx context.Context, in *npool.GetNAppOpLogsRequest) (*npool.GetNAppOpLogsResponse, error) {
	handler, err := oplog1.NewHandler(
		ctx,
		oplog1.WithAppID(in.GetTargetAppID(), true),
		oplog1.WithOffset(in.GetOffset()),
		oplog1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNAppOpLogs",
			"In", in,
			"Error", err,
		)
		return &npool.GetNAppOpLogsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetOpLogs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetNAppOpLogs",
			"In", in,
			"Error", err,
		)
		return &npool.GetNAppOpLogsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &npool.GetNAppOpLogsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
