package utilits

import (
	"context"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const MAX_GRPC_SIZE = 1024 * 1024 * 100

type ExpectedConnections struct {
	SqlConnection         *sqlx.DB
	SessionGrpcConnection *grpc.ClientConn
	UtilsRedisPool        *redis.Pool
}

func NewGrpcConnection(grpcUrl string) (*grpc.ClientConn, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	return grpc.DialContext(ctx, grpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(MAX_GRPC_SIZE),
			grpc.MaxCallSendMsgSize(MAX_GRPC_SIZE)),
		grpc.WithBlock())
}
