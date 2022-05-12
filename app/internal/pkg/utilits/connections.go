package utilits

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

const MAX_GRPC_SIZE = 1024 * 1024 * 100

type ExpectedConnections struct {
	SqlConnection         *sqlx.DB
	SessionGrpcConnection *grpc.ClientConn
	UtilsRedisPool        *redis.Pool
}

func NewGrpcConnection(grpcUrl string) (*grpc.ClientConn, error) {
	return grpc.Dial(grpcUrl, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(MAX_GRPC_SIZE),
		grpc.MaxCallSendMsgSize(MAX_GRPC_SIZE)), grpc.WithBlock())
}
