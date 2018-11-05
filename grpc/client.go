package grpc

import (
	"github.com/KAMESTERY/middlewarekam/utils"
	"google.golang.org/grpc"
)

var client_logger = utils.NewLogger("grpcclient")

func NewGrpcConn() (grpcConn *grpc.ClientConn, err error) {
	grpcConn, err = grpc.Dial(utils.BackendRPC, grpc.WithInsecure())
	if err != nil {
		client_logger.Fatalf("GRPC_CONNECTION_ERROR:::: %v", err)
	}
	return
}
