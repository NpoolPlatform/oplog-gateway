package oplog

import (
	"context"

	oplog1 "github.com/NpoolPlatform/message/npool/oplog/gw/v1/oplog"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	oplog1.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	oplog1.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := oplog1.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
