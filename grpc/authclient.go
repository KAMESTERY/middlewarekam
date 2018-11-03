package grpc

import (
	"github.com/outcastgeek/kamworks/utils"
	"google.golang.org/grpc"
)

var authclient_logger = utils.NewLogger("entpointauthclient")

func NewGrpcConn() (grpcConn *grpc.ClientConn, err error) {
	grpcConn, err = grpc.Dial(utils.BackendRPC, grpc.WithInsecure())
	if err != nil {
		authclient_logger.Fatalf("GRPC_CONNECTION_ERROR:::: %v", err)
	}
	return
}
