package api

import (
	"context"

	servicetmpl "github.com/NpoolPlatform/message/npool/oplog/gw/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	servicetmpl.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	servicetmpl.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := servicetmpl.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
